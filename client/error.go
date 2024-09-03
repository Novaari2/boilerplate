package client

import "fmt"

type Error struct {
	Message    string
	StatusCode int
	RawError   error
}

func (e *Error) Error() string {
	if e.RawError != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.RawError.Error)
	}

	return e.Message
}

func (e *Error) Unwrap() error {
	return e.RawError
}
