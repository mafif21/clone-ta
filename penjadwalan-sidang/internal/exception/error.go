package exception

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (error *ErrorResponse) Error() string {
	return error.Message
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}
