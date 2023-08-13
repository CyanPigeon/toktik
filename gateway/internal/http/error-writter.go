package http

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type defaultErrorWriter struct{}

func NewErrorWriter() ErrorWriter {
	return &defaultErrorWriter{}
}

func (t *defaultErrorWriter) Write(writer http.ResponseWriter, status int, err error) {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	json, e := json2.Marshal(struct {
		StatusCode int    `json:"status_code,omitempty"`
		StatusMsg  string `json:"status_msg,omitempty"`
	}{
		StatusCode: 301,
		StatusMsg:  err.Error(),
	})
	if e != nil {
		panic(e)
	}
	if _, e = writer.Write(json); e != nil {
		panic(fmt.Errorf("unexpected error when processing error: %+v", e))
	}
}
