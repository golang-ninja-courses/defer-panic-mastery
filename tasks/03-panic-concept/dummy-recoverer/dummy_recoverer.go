package recoverer

type RecoveryHandler func(err interface{})

// DummyRecoverer предполагается для использования в тестах.
// Он не ловит паники в принципе.
type DummyRecoverer struct{}

func NewDummyRecoverer() DummyRecoverer {
	// Реализуй меня.
	return DummyRecoverer{}
}

func (r DummyRecoverer) Do(f func()) {
	// Реализуй меня.
}

func (r DummyRecoverer) DoWithRecoveryHandler(f func(), handler RecoveryHandler) {
	// Реализуй меня.
}

func (r DummyRecoverer) Go(f func()) {
	// Реализуй меня.
}

func (r DummyRecoverer) GoWithRecoveryHandler(f func(), handler RecoveryHandler) {
	// Реализуй меня.
}
