package errors

type StatusResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message,omitempty"`
	Error   *ExportableError `json:"error,omitempty"`
}

func NewSuccessResponse() *StatusResponse {
	return &StatusResponse{
		Status: "success",
	}
}

func NewErrorResponse(err *ExportableError) *StatusResponse {
	return &StatusResponse{
		Status: "error",
		Error:  err,
	}
}
