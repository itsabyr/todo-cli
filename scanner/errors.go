package scanner

import "errors"

var (
	errEmptyInput = errors.New("empty input")
	errNeedExit   = errors.New("exit requested")
	errWrongArgs  = errors.New("wrong number of arguments")
	errUnknownCmd = errors.New("unknown command")
)
