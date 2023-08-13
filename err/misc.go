package err

import (
	"errors"
	"github.com/CyanPigeon/toktik/util"
)

type ErrInvalidRequest struct{}

func (e *ErrInvalidRequest) Error() error {
	return errors.New("")
}

func (e *ErrInvalidRequest) ErrorJson() string {
	return util.GenError(e)
}

func (e *ErrInvalidRequest) Code() int {
	return 301
}

func (e *ErrInvalidRequest) Ok() bool {
	return false
}

type ErrUnknown struct{}

func (e *ErrUnknown) Error() error {
	return errors.New("")
}

func (e *ErrUnknown) ErrorJson() string {
	return util.GenError(e)
}

func (e *ErrUnknown) Code() int {
	return 999
}

func (e *ErrUnknown) Ok() bool {
	return false
}
