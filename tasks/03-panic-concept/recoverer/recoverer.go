package recoverer

type RecoveryHandler func(err any)

type GoroutineStarter interface {
	Go(func())
}

type Recoverer struct {
	handler RecoveryHandler
	starter GoroutineStarter
}

// NewRecoverer создаёт новый Recoverer, в качестве аргументов принимает
// обработчик паники по умолчанию и инструмент для запуска горутин.
func NewRecoverer(h RecoveryHandler, s GoroutineStarter) *Recoverer {
	// Реализуй меня.
	return nil
}

// Do синхронно выполняет функцию f.
// Возможная паника обрабатывается хендлером по умолчанию.
func (r *Recoverer) Do(f func()) {
	// Реализуй меня.
}

// DoWithRecoveryHandler синхронно выполняет функцию f.
// Возможная паника обрабатывается хендлером handler.
func (r *Recoverer) DoWithRecoveryHandler(f func(), handler RecoveryHandler) {
	// Реализуй меня.
}

// Go выполняет функцию f в новой горутине, созданной с помощью GoroutineStarter.
// Возможная паника обрабатывается хендлером по умолчанию.
func (r *Recoverer) Go(f func()) {
	// Реализуй меня.
}

// GoWithRecoveryHandler выполняет функцию f в новой горутине, созданной с помощью GoroutineStarter.
// Возможная паника обрабатывается хендлером handler.
func (r *Recoverer) GoWithRecoveryHandler(f func(), handler RecoveryHandler) {
	// Реализуй меня.
}
