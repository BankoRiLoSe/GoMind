package dto

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(data any) Response {
	return Response{
		Code:    0,
		Message: "ok",
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
