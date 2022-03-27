package panichelpers

// Recover вызывает fn, ловит панику и возвращает её аргумент и true.
// Если паники не было, то функция вторым аргументом вернёт false.
func Recover(fn func()) (interface{}, bool) {
	// Реализуй меня.
	return nil, false
}
