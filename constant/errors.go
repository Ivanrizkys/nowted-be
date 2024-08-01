package constant

import "errors"

type errorWrapper struct {
	err error
}

func (e *errorWrapper) Error() string {
	return e.err.Error()
}

var (
	ErrNotFound   = &errorWrapper{errors.New("not found error")}
	ErrBadRequest = &errorWrapper{errors.New("bad request error")}
	ErrService    = &errorWrapper{errors.New("service error")}
	ErrUnAuth     = &errorWrapper{errors.New("not authorized")}
)
