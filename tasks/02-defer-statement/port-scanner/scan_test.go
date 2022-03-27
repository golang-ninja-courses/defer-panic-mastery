package portscanner

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestScan_Example(t *testing.T) {
	avPorts := Scan(realDialer{}, "localhost")
	t.Log(avPorts)

	// При верной реализации вы увидите порты, открытые в данный момент на вашей системе:
	// [631 3306 6942 33060 55412 63342]
}

func TestScan_Smoke(t *testing.T) {
	const host = "www.defense.gov"

	presets := make(map[portNum]*connMock)
	d := newDialerMock(t, host, presets)
	for p := 1; p <= 65535; p++ {
		presets[p] = newConnMock(d, time.Microsecond, syscall.ECONNREFUSED)
	}
	for _, p := range []int{631, 3306, 6942, 33060, 55412, 63342} {
		presets[p] = newConnMock(d, time.Microsecond, nil)
	}

	avPorts := Scan(d, host)
	d.assertAllPortsAreVisited()
	d.assertNoConnectionsLeak()
	assert.Equal(t, []int{631, 3306, 6942, 33060, 55412, 63342}, avPorts)
}

func TestScan_ConcurrencyAndHoneypots(t *testing.T) {
	const host = "www.defense.gov"

	presets := make(map[portNum]*connMock)
	d := newDialerMock(t, host, presets)
	for p := 1; p <= 65535; p++ {
		presets[p] = newConnMock(d, time.Microsecond, syscall.ECONNREFUSED)
	}
	for i := 0; i < 20; i++ {
		presets[rand.Intn(65536)+1] = newConnMock(d, time.Second, nil)
	}
	for _, i := range []int{5423, 8080} {
		presets[i] = newConnMock(d, time.Microsecond, nil)
	}

	start := time.Now()
	avPorts := Scan(d, host)
	elapsed := time.Since(start)

	assert.Less(t, elapsed, 2*time.Second, "are ports scanned concurrently?")
	d.assertAllPortsAreVisited()
	d.assertNoConnectionsLeak()
	assert.Equal(t, []int{5423, 8080}, avPorts)
}

func TestScan_PortsSorting(t *testing.T) {
	const host = "www.defense.gov"

	presets := make(map[portNum]*connMock)
	d := newDialerMock(t, host, presets)
	for p := 1; p <= 65535; p++ {
		presets[p] = newConnMock(d, time.Microsecond, nil)
	}

	avPorts := Scan(d, host)
	d.assertAllPortsAreVisited()
	d.assertNoConnectionsLeak()

	expected := make([]int, 65535)
	for i := 0; i < 65535; i++ {
		expected[i] = i + 1
	}
	assert.Equal(t, expected, avPorts)
}

type realDialer struct{}

func (d realDialer) DialContext(ctx context.Context, network, address string) (Conn, error) {
	return new(net.Dialer).DialContext(ctx, network, address)
}

type portNum = int

type dialerMock struct {
	t               *testing.T
	openConnections int
	portVisited     [65535 + 1]bool
	// mu защищает поля выше.
	mu sync.Mutex

	host    string
	presets map[portNum]*connMock
}

func newDialerMock(t *testing.T, host string, presets map[portNum]*connMock) *dialerMock { //nolint:thelper
	d := &dialerMock{
		t:       t,
		host:    host,
		presets: presets,
	}
	return d
}

type connMock struct {
	establishTime time.Duration
	err           error
	onClose       func()
}

func newConnMock(d *dialerMock, t time.Duration, err error) *connMock {
	return &connMock{
		err:           err,
		establishTime: t,
		onClose: func() {
			d.mu.Lock()
			d.openConnections--
			d.mu.Unlock()
		},
	}
}

func (c *connMock) Close() error {
	c.onClose()
	return nil
}

func (d *dialerMock) DialContext(ctx context.Context, network, address string) (conn Conn, err error) {
	if network != "tcp" {
		return nil, fmt.Errorf("unsupported network: %s", network)
	}

	h, p, err := splitHostPort(address)
	if err != nil {
		return nil, err
	}
	if h != d.host {
		return nil, fmt.Errorf("unexpected hostname: %s", h)
	}
	if p < 1 || p > 65535 {
		return nil, fmt.Errorf("unexpected port: %d", p)
	}

	c, ok := d.presets[p]
	if !ok {
		panic(fmt.Sprintf("no preset for port %d", p))
	}

	func() {
		d.mu.Lock()
		defer d.mu.Unlock()

		d.openConnections++
		if d.openConnections > 10 {
			panic("too many open connections")
		}
	}()

	defer func() {
		// Если возвращаем ошибку, то сами эмулируем закрытие соединения,
		// иначе ожидается вызов conn.Close().
		if err != nil {
			d.mu.Lock()
			d.openConnections--
			d.mu.Unlock()
		}
	}()

	d.mu.Lock()
	d.portVisited[p] = true
	d.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(c.establishTime):
	}

	if c.err != nil {
		return nil, fmt.Errorf("connect: %w", c.err)
	}
	return c, nil
}

func (d *dialerMock) assertAllPortsAreVisited() {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i := 1; i <= 65535; i++ {
		assert.True(d.t, d.portVisited[i], "port %d was not scanned", i)
	}
}

func (d *dialerMock) assertNoConnectionsLeak() {
	d.mu.Lock()
	defer d.mu.Unlock()

	assert.Equal(d.t, 0, d.openConnections, "connections leak")
}

func splitHostPort(addr string) (string, int, error) {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("%s: invalid addr format: `host:port` expected", addr)
	}

	p, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("%s: port is not a number", parts[1])
	}

	return parts[0], p, nil
}
