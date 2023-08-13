package err

type IError interface {
	Error() error
	ErrorJson() string
	Code() int
	Ok() bool
}
