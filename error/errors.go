package error

import "fmt"

type Http struct {
	ErrorMessage  string `json:"error_message,omitempty"`
	ErrorMetadata string `json:"error_metadata,omitempty"`
	ErrorCode     int    `json:"error_code"`
}

func (e Http) Error() string {
	return fmt.Sprintf("description: %s,  metadata: %s", e.ErrorMessage, e.ErrorMetadata)
}

func NewHttpError(errorMessage, errorMetadata string, errorCode int) *Http {
	return &Http{
		ErrorMessage:  errorMessage,
		ErrorMetadata: errorMetadata,
		ErrorCode:     errorCode,
	}
}
