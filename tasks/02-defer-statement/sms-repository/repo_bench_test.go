package smsrepository

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

/*
Для авторского решения:

$ go test -benchmem -bench .
BenchmarkRepo-8              146           8172592 ns/op             175 B/op          1 allocs/op
PASS
ok      tasks/02-defer-statement/sms-repository      2.076s
*/

func BenchmarkRepo(b *testing.B) {
	const (
		messages     = 100
		opsPerWorker = 10_000
	)

	r := NewRepo()

	getRandomMsgID := func() MessageID {
		return MessageID(strconv.Itoa(rand.Intn(messages)))
	}

	saver := func() {
		for i := 0; i < opsPerWorker; i++ {
			_ = r.Save(getRandomMsgID())
		}
	}

	getter := func() {
		for i := 0; i < opsPerWorker; i++ {
			_, _ = r.Get(getRandomMsgID())
		}
	}

	updater := func() {
		for i := 0; i < opsPerWorker; i++ {
			id := getRandomMsgID()
			s, err := r.Get(id)
			if err != nil {
				return
			}

			var newStatus MessageStatus
			switch s {
			case MessageStatusAccepted:
				newStatus = MessageStatusConfirmed
			case MessageStatusConfirmed:
				newStatus = []MessageStatus{MessageStatusDelivered, MessageStatusFailed}[rand.Intn(2)]
			case MessageStatusFailed:
				newStatus = MessageStatusFailed
			case MessageStatusDelivered:
				newStatus = MessageStatusDelivered
			}

			_ = r.Update(id, newStatus)
		}
	}

	ops := []func(){
		getter,
		saver,
		updater,
	}

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(len(ops))

		for _, op := range ops {
			op := op
			go func() {
				defer wg.Done()
				op()
			}()
		}

		wg.Wait()
	}
}
