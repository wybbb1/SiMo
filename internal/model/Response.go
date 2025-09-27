package model

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(data interface{}) *Response {
	return &Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}	

func Error(message string) *Response {
	return &Response{
		Code:    500,
		Message: message,
	}
}

func ErrorWithCode(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}

