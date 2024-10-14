package core

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "time"
)

type Error struct {
	Code int
	Message string
	OccurredAt time.Time
	OriginalErr error
}

func NewError(code int, msg string, origin ...error) *Error {
	var originErr error
	if len(origin) > 0 {
		originErr = origin[0]
	}

	return &Error{
		Code: code,
		Message: msg,
		OccurredAt: time.Now(),
		OriginalErr: originErr,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error %d: %s (Occurred at: %s) with origin from: %v", e.Code, e.Message, e.OccurredAt.Format(time.RFC3339), e.OriginalErr)
}

func (e *Error) Respons() gin.H {
	return gin.H{
		"code": e.Code,
		"message": e.Message,
		"occurredAt": e.OccurredAt,
		"originalErr": e.OriginalErr,
	}
}