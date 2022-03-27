package fnhelpers

type ExitReason int

const (
	// ExitReasonRegularReturn означает, что функция завершилась в штатном режиме.
	ExitReasonRegularReturn ExitReason = iota + 1

	// ExitReasonPanic означает, что функция запаниковала.
	ExitReasonPanic

	// ExitReasonGoexit означает, что функция вызвала runtime.Goexit.
	ExitReasonGoexit
)

// WhatHappened вызывает fn и сообщает о статусе её завершения.
func WhatHappened(fn func()) ExitReason {
	return 0
}
