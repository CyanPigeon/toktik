package errors

type RequestInterruptError struct {
	Cause  error
	Status int
}

func (t *RequestInterruptError) Error() string {
	return t.Cause.Error()
}
