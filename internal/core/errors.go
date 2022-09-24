package core

import (
	"fmt"
	"github.com/pkg/errors"
)

type Error struct {
	Cause string
	trace error
	Msg   string
}

func (e *Error) Error() string {
	return fmt.Sprintf("[message] %s --- [cause] %s --- [trace] %s ", e.Msg, e.Cause, e.trace)
}

func (e *Error) SetCause(cause string) *Error {
	if e.trace == nil {
		return &Error{Cause: cause, trace: errors.New(cause), Msg: e.Msg}
	}

	return &Error{Cause: cause, trace: e.trace, Msg: e.Msg}
}

func (e *Error) WrapTrace(trace string) *Error {
	return &Error{Cause: e.Cause, trace: errors.Wrap(e.trace, trace), Msg: e.Msg}
}

func NewError(cause, msg string, trace error) *Error {
	return &Error{
		Cause: cause,
		trace: trace,
		Msg:   msg,
	}
}
