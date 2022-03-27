package example

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("1", func(tt *testing.T) {
		tt.Run("1.1", func(ttt *testing.T) { fmt.Println("1.1") })
		tt.Run("1.2", func(ttt *testing.T) { fmt.Println("1.2"); tt.FailNow() }) // Не ttt!
		fmt.Println("1")
	})

	t.Run("2", func(tt *testing.T) {
		tt.Run("2.1", func(ttt *testing.T) { fmt.Println("2.1") })
		tt.Run("2.2", func(ttt *testing.T) { fmt.Println("2.2") })
		fmt.Println("2")
	})
}

/*
$ go test  .
1.1
1.2
2.1
2.2
2
--- FAIL: TestExample (0.00s)
    --- FAIL: TestExample/1 (0.00s)
        --- FAIL: TestExample/1/1.2 (0.00s)
            testing.go:1169: test executed panic(nil) or runtime.Goexit: subtest may have called FailNow on a parent test
*/
