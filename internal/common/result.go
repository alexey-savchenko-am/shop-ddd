package common

type Result[T any] struct {
	IsSuccess bool   `json:"is_success"`
	Value     T      `json:"value,omitempty"`
	Error     *Error `json:"error,omitempty"`
}

func Success[T any](value T) Result[T] {
	return Result[T]{IsSuccess: true, Value: value}
}

func Failure[T any](err Error) Result[T] {
	return Result[T]{IsSuccess: false, Error: &err}
}

type EmptyResult = Result[struct{}]

func SuccessEmpty() EmptyResult {
	return Result[struct{}]{IsSuccess: true}
}

func FailureEmpty(err Error) EmptyResult {
	return Result[struct{}]{IsSuccess: false, Error: &err}
}
