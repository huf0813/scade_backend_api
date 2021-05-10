package custom_response

type CustomResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewCustomResponse(success bool, message string, data interface{}) CustomResponse {
	return CustomResponse{
		Data:    data,
		Message: message,
		Success: success,
	}
}
