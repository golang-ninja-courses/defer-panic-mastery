//go:build NDEBUG

package assert

func Assert(cond bool, msg string, args ...any) {}
