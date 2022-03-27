package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

//go:generate go run $GOFILE
// go run -tags test main.go

const n = 5_000_000

var t = template.Must(template.New("defers").Parse(`// +build test
package main

func main() {
{{ .Defers -}}
}

func foo() {}
`))

func main() {
	f, err := os.Create("main.go")
	mustNil(err)
	defer f.Close()

	mustNil(t.Execute(f, struct {
		Defers string
	}{
		Defers: strings.Repeat("\tdefer foo()\n", n),
	}))
	mustNil(f.Sync())
}

func mustNil(err error) {
	if err != nil {
		log.Panic(err)
	}
}
