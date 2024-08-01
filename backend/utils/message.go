package utils

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}) Response {
	return Response{
		Code:    0,
		Message: message,
		Data:    data,
	}
}

func Error(message string) Response {
	return Response{
		Code:    1,
		Message: message,
		Data:    nil,
	}
}
