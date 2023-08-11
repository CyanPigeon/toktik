package err

import "reflect"

type ErrUserNotExist struct {
	msg string
}

func (e ErrUserNotExist) Error() string {
	return "user not exist"
}

func (e ErrUserNotExist) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}

type ErrUserActionFailed struct {
	msg string
}

func (e ErrUserActionFailed) Error() string {
	return "user action failed"
}

func (e ErrUserActionFailed) Is(err error) bool {
	return reflect.TypeOf(err).Name() == reflect.TypeOf(e).Name()
}
