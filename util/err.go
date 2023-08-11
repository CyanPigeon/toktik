package util

import (
	"encoding/json"
	"errors"
	e "playground/err"
)

type ErrorJson struct {
	StatusCode int    `json:"status_code,omitempty"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

const (
	OK                  = 0
	Unknown             = 999
	LoginRegisterFailed = 101
	Unauthorized        = 102
	UserNotExist        = 201
	UserActionFailed    = 202
	InvalidRequest      = 301
)

func genErrorStruct(code int, err error) ErrorJson {
	return ErrorJson{
		StatusCode: code,
		StatusMsg:  err.Error(),
	}
}

func GenErrorJson(err error) string {
	if errors.Is(err, e.ErrAuthFailed{}) {
		j, _ := json.Marshal(
			genErrorStruct(LoginRegisterFailed, err),
		)
		return string(j)
	}
	if errors.Is(err, e.ErrInvalidAccount{}) {
		j, _ := json.Marshal(
			genErrorStruct(LoginRegisterFailed, err),
		)
		return string(j)
	}
	if errors.Is(err, e.ErrBadRequest{}) {
		j, _ := json.Marshal(
			genErrorStruct(InvalidRequest, err),
		)
		return string(j)
	}
	if errors.Is(err, e.ErrUserNotExist{}) {
		j, _ := json.Marshal(
			genErrorStruct(UserNotExist, err),
		)
		return string(j)
	}
	if errors.Is(err, e.ErrUnauthorized{}) {
		j, _ := json.Marshal(
			genErrorStruct(Unauthorized, err),
		)
		return string(j)
	}
	if errors.Is(err, e.ErrUserActionFailed{}) {
		j, _ := json.Marshal(
			genErrorStruct(UserActionFailed, err),
		)
		return string(j)
	}

	j, _ := json.Marshal(
		genErrorStruct(Unknown, err),
	)
	return string(j)
}
