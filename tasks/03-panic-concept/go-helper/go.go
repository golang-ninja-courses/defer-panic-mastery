package errhelpers

import "errors"

var ErrPanicOccurred = errors.New("panic occurred")

// Go позволяет запустить функцию fn в отдельной горутине
// и получить результат её работы через выходной канал.
// Если функция запаникует, то в канале окажется ErrPanicOccurred.
func Go(fn func() error) <-chan error {
	// Реализуй меня.
	return nil
}
