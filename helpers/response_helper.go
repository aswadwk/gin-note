package helpers

type Response struct {
	Status  string        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyResponse struct{}

func ResponseSuccess(status string, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}

func ResponseError(status string, message string, data interface{}) Response {

	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return res
}