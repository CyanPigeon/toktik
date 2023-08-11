package err

import (
	"reflect"
)

type ErrUnknown struct {
	msg string
}

func (e ErrUnknown) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}

func (e ErrUnknown) Error() string {
	return "unknown error"
}

type ErrBadRequest struct {
	msg string
}

func (e ErrBadRequest) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}

func (e ErrBadRequest) Error() string {
	return "unknown error"
}
