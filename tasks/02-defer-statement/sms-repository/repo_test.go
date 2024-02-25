package smsrepository

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestRepo_Save(t *testing.T) {
	t.Run("success save", func(t *testing.T) {
		r := NewRepo()

		for _, id := range []MessageID{"1", "2", "3"} {
			err := r.Save(id)
			require.NoError(t, err, "id = %q", id)

			s, err := r.Get(id)
			require.NoError(t, err, "id = %q", id)
			assert.Equal(t, MessageStatusAccepted, s, "id = %q", id)
		}
	})

	t.Run("msg already exist", func(t *testing.T) {
		r := NewRepo()
		msgs := []MessageID{"1", "2", "3"}

		for _, id := range msgs {
			err := r.Save(id)
			require.NoError(t, err, "id = %q", id)
		}

		for _, id := range msgs {
			err := r.Save(id)
			require.ErrorIs(t, err, ErrMsgAlreadyExists, "id = %q", id)
		}

		err := r.Save("4")
		require.NoError(t, err)
	})
}

func TestRepo_Get_NonExistentMsg(t *testing.T) {
	r := NewRepo()

	for _, id := range []MessageID{"1", "2", "3"} {
		s, err := r.Get(id)
		require.ErrorIs(t, err, ErrMsgNotFound, "id = %q", id)
		assert.Equal(t, MessageStatus(""), s, "id = %q", id)
	}
}

func TestRepo_Update(t *testing.T) {
	t.Run("non-existent msg", func(t *testing.T) {
		r := NewRepo()

		for _, id := range []MessageID{"1", "2", "3"} {
			err := r.Update(id, MessageStatusConfirmed)
			require.ErrorIs(t, err, ErrMsgNotFound, "id = %q", id)
		}
	})

	t.Run("status paths", func(t *testing.T) {
		r := NewRepo()

		assertStatus := func(t *testing.T, id MessageID, expected MessageStatus) {
			t.Helper()

			s, err := r.Get(id)
			require.NoError(t, err)
			assert.Equal(t, expected, s)
		}

		updateWithoutErr := func(t *testing.T, id MessageID, newStatus MessageStatus) {
			t.Helper()

			err := r.Update(id, newStatus)
			require.NoError(t, err)

			assertStatus(t, id, newStatus)
		}

		updateWithErr := func(t *testing.T, id MessageID, newStatus MessageStatus) {
			t.Helper()

			prevStatus, err := r.Get(id)
			require.NoError(t, err)

			err = r.Update(id, newStatus)
			require.ErrorIs(t, err, ErrInvalidMsgStatusChange)

			assertStatus(t, id, prevStatus)
		}

		const (
			msgID1 MessageID = "1"
			msgID2 MessageID = "2"
			msgID3 MessageID = "3"
		)

		t.Run("accepted confirmed", func(t *testing.T) {
			err := r.Save(msgID1)
			require.NoError(t, err)

			updateWithoutErr(t, msgID1, MessageStatusAccepted) // Accepted -> Accepted
			for _, s := range []MessageStatus{MessageStatusFailed, MessageStatusDelivered} {
				updateWithErr(t, msgID1, s)
			}
			updateWithoutErr(t, msgID1, MessageStatusConfirmed) // Accepted -> Confirmed
			updateWithoutErr(t, msgID1, MessageStatusConfirmed) // Confirmed -> Confirmed
			updateWithErr(t, msgID1, MessageStatusAccepted)
		})

		t.Run("confirmed failed", func(t *testing.T) {
			err := r.Save(msgID2)
			require.NoError(t, err)

			updateWithoutErr(t, msgID2, MessageStatusConfirmed) // Accepted -> Confirmed
			updateWithoutErr(t, msgID2, MessageStatusFailed)    // Confirmed -> Failed
			for _, s := range []MessageStatus{MessageStatusAccepted, MessageStatusConfirmed, MessageStatusDelivered} {
				updateWithErr(t, msgID2, s)
			}
			updateWithoutErr(t, msgID2, MessageStatusFailed) // Failed -> Failed
		})

		t.Run("confirmed delivered", func(t *testing.T) {
			err := r.Save(msgID3)
			require.NoError(t, err)

			updateWithoutErr(t, msgID3, MessageStatusConfirmed) // Accepted -> Confirmed
			updateWithoutErr(t, msgID3, MessageStatusDelivered) // Confirmed -> Delivered
			for _, s := range []MessageStatus{MessageStatusAccepted, MessageStatusConfirmed, MessageStatusFailed} {
				updateWithErr(t, msgID3, s)
			}
			updateWithoutErr(t, msgID3, MessageStatusDelivered) // Delivered -> Delivered
		})

		assertStatus(t, msgID1, MessageStatusConfirmed)
		assertStatus(t, msgID2, MessageStatusFailed)
		assertStatus(t, msgID3, MessageStatusDelivered)
	})
}

func TestRepo_ConcurrentAccess(t *testing.T) {
	const (
		workers      = 10
		opsPerWorker = 10_000
		messages     = 100
	)

	r := NewRepo()

	getRandomMsgID := func() MessageID {
		return MessageID(strconv.Itoa(rand.Intn(messages)))
	}

	getRandomMsgStatus := func() MessageStatus {
		statuses := []MessageStatus{
			MessageStatusAccepted,
			MessageStatusConfirmed,
			MessageStatusFailed,
			MessageStatusDelivered,
		}
		return statuses[rand.Intn(len(statuses))]
	}

	saveRandomMsg := func() {
		_ = r.Save(getRandomMsgID())
	}

	getRandomMsg := func() {
		_, _ = r.Get(getRandomMsgID())
	}

	updateRandomMsgWithRandomStatus := func() {
		_ = r.Update(getRandomMsgID(), getRandomMsgStatus())
	}

	getRandomOp := func() func() {
		ops := []func(){
			saveRandomMsg,
			getRandomMsg,
			updateRandomMsgWithRandomStatus,
		}
		return ops[rand.Intn(len(ops))]
	}

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < opsPerWorker; j++ {
				getRandomOp()()
			}
		}()
	}

	wg.Wait()
}

func TestRepo_PartialLock(t *testing.T) {
	const msgID = "1"

	tryToHack := func() (hacked bool) {
		r := NewRepo()
		require.NoError(t, r.Save(msgID))

		err := r.Update(msgID, MessageStatusConfirmed)
		require.NoError(t, err)

		var wg sync.WaitGroup
		wg.Add(2)
		var updErrors [2]error
		go func() { updErrors[0] = r.Update(msgID, MessageStatusDelivered); wg.Done() }()
		go func() { updErrors[1] = r.Update(msgID, MessageStatusFailed); wg.Done() }()
		wg.Wait()

		s, err := r.Get(msgID)
		require.NoError(t, err)
		assert.Contains(t, []MessageStatus{MessageStatusDelivered, MessageStatusFailed}, s)

		// Оба Update смогли "проскочить" и завершились без ошибки.
		return (updErrors[0] == nil) && (updErrors[1] == nil)
	}

	for i := 0; i < 10_000; i++ {
		if hacked := tryToHack(); hacked {
			t.Fatalf(`Update() method contains "concurrent holes" (iteration=%d)`, i)
		}
	}
}
