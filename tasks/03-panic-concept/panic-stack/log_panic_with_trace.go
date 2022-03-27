package panichelpers

type Logger interface {
	Logf(format string, args ...interface{})
}

func LogPanicWithTrace(l Logger) {
	// Реализуй меня.
}
