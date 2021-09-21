package response

type ActionResultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewSuccessResponse(message string) ActionResultResponse {
	return ActionResultResponse{
		Success: true,
		Message: message,
	}
}

func NewErrorResponse(message string) ActionResultResponse {
	return ActionResultResponse{
		Success: false,
		Message: message,
	}
}
