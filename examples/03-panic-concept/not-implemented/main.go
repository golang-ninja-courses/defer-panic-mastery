package main

import "errors"

var ErrNotImplemented = errors.New("not implemented")

type Writer struct{}

func (w Writer) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (w Writer) Close() error {
	return ErrNotImplemented
}
