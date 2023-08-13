package util

import (
	"encoding/json"
	"github.com/CyanPigeon/toktik/err"
)

type ErrorStruct struct {
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

func genErrorStruct(err err.IError) ErrorStruct {
	code := err.Code()
	switch code {
	case LoginRegisterFailed:
		return ErrorStruct{
			StatusCode: LoginRegisterFailed,
			StatusMsg:  "login or register failed",
		}
	case Unauthorized:
		return ErrorStruct{
			StatusCode: Unauthorized,
			StatusMsg:  "unauthorized",
		}
	case UserNotExist:
		return ErrorStruct{
			StatusCode: UserNotExist,
			StatusMsg:  "user not exist",
		}
	case UserActionFailed:
		return ErrorStruct{
			StatusCode: UserActionFailed,
			StatusMsg:  "user action failed",
		}
	case InvalidRequest:
		return ErrorStruct{
			StatusCode: InvalidRequest,
			StatusMsg:  "invalid request",
		}
	default:
		return ErrorStruct{
			StatusCode: Unknown,
			StatusMsg:  "unknown",
		}
	}
}

func GenError(err err.IError) string {
	if err != nil {
		j, _ := json.Marshal(genErrorStruct(err))
		return string(j)
	}
	return ""
}
