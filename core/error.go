package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code        int
	Message     string
	OriginalErr error
}

func NewError(code int, msg string, origin ...error) *Error {
	var originErr error
	if len(origin) > 0 {
		originErr = origin[0]
	}

	return &Error{
		Code:        code,
		Message:     msg,
		OriginalErr: originErr,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %d: %s (Occurred at: %s) with origin from: %v", e.Code, e.Message, e.OriginalErr)
}

func (e *Error) Respons() gin.H {
	return gin.H{
		"code":        e.Code,
		"message":     e.Message,
		"originalErr": e.OriginalErr,
	}
}
