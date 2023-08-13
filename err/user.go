package err

import (
	"errors"
	"github.com/CyanPigeon/toktik/util"
)

type ErrUserLoginOrRegister struct{}

func (e *ErrUserLoginOrRegister) Error() error {
	return errors.New("")
}

func (e *ErrUserLoginOrRegister) ErrorJson() string {
	return util.GenError(e)
}

func (e *ErrUserLoginOrRegister) Code() int {
	return 101
}

func (e *ErrUserLoginOrRegister) Ok() bool {
	return false
}

type ErrUserUnauth struct{}

func (e *ErrUserUnauth) Error() error {
	return errors.New("")
}

func (e *ErrUserUnauth) ErrorJson() string {
	return util.GenError(e)
}

func (e *ErrUserUnauth) Code() int {
	return 102
}

func (e *ErrUserUnauth) Ok() bool {
	return false
}

type ErrUserNotExist struct{}

func (e *ErrUserNotExist) Error() error {
	return errors.New("")
}

func (e *ErrUserNotExist) ErrorJson() string {
	return util.GenError(e)
}

func (e *ErrUserNotExist) Code() int {
	return 201
}

func (e *ErrUserNotExist) Ok() bool {
	return false
}

type ErrUserAction struct{}

func (e *ErrUserAction) Error() error {
	return errors.New("")
}

func (e *ErrUserAction) ErrorJson() string {
	return util.GenError(e)
}

func (e *ErrUserAction) Code() int {
	return 202
}

func (e *ErrUserAction) Ok() bool {
	return false
}
