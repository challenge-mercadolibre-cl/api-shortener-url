package exceptions

import "errors"

var (
	ErrUnexpectedCommand = errors.New("unexpected command")
)
