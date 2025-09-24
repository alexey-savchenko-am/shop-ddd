package common

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Inner   *Error `json:"inner,omitempty"`
}

var None = Error{}

func (e Error) Error() string {
	if e.Inner != nil {
		return e.Message + " -> " + e.Inner.Error()
	}
	return e.Message
}

func NewError(code, message string) Error {
	return Error{Code: code, Message: message}
}

func WrapError(err Error, code, message string) Error {
	return Error{
		Code:    code,
		Message: message,
		Inner:   &err,
	}
}

func FromError(code string, err error) Error {

	if err == nil {
		return None
	}

	if e, ok := err.(Error); ok {
		return e
	}

	if e, ok := err.(*Error); ok {
		return *e
	}

	return Error{
		Code:    code,
		Message: err.Error(),
	}
}
