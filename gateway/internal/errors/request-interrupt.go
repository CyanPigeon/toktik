package errors

type RequestInterruptError struct {
	cause  error
	status int
}

func (t *RequestInterruptError) Error() string {
	return t.cause.Error()
}

func NewRequestInterruptError(status int, cause error) *RequestInterruptError {
	return &RequestInterruptError{
		cause:  cause,
		status: status,
	}
}

func (t *RequestInterruptError) Cause() error {
	return t.cause
}

func (t *RequestInterruptError) Status() int {
	return t.status
}
