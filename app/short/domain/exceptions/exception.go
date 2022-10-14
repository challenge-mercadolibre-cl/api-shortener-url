package exceptions

import "errors"

var (
	ErrUrlIdEmpty  = errors.New("url id cant be empty")
	ErrUserIdEmpty = errors.New("user id cant be empty")
	ErrUrlFormat   = errors.New("format url invalid")
)
