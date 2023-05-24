package errors

type WrapError struct {
	msg string
}

func (e WrapError) Error() string {
	return e.msg
}

func NewWrapError(e string) *WrapError {
	return &WrapError{msg: e}
}
