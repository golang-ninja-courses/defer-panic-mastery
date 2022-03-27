package deferr

type Defer struct{}

func NewDefer() *Defer {
	// Реализуй меня.
	return nil
}

// Defer добавляет функцию в стек отложенных функций.
func (d *Defer) Defer(f func()) {
	// Реализуй меня.
}

// Execute выполняет функции, добавленные через Defer, в порядке, обратном порядку добавления.
// Предполагается, что Execute вызывается единожды и после него работа с объектом типа *Defer заканчивается.
func (d *Defer) Execute() {
	// Реализуй меня.
}
