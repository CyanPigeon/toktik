package http

import (
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

type defaultErrorWriter struct{}

func NewErrorWriter() ErrorWriter {
	return &defaultErrorWriter{}
}

func (t *defaultErrorWriter) Write(writer http.ResponseWriter, status int, err error) {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, e := writer.Write([]byte(err.Error())); e != nil {
		log.Errorf("unexpected error when processing error: %v", e)
	}
}
