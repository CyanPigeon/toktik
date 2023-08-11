package err

import (
	"reflect"
)

type ErrAuthFailed struct {
	msg string
}

func (e ErrAuthFailed) Error() string {
	return "auth failed"
}

func (e ErrAuthFailed) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}

type ErrInvalidAccount struct {
	msg string
}

func (e ErrInvalidAccount) Error() string {
	return "invalid account"
}

func (e ErrInvalidAccount) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}

type ErrUnauthorized struct {
	msg string
}

func (e ErrUnauthorized) Error() string {
	return "unauthorized"
}

func (e ErrUnauthorized) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}
