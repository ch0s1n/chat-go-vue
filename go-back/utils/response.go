package utils

type Ecode struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HandleResponse(message string, data any) Ecode {
	return Ecode{
		Message: message,
		Data:    data,
	}
}
