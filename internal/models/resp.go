package models

import "github.com/akram620/alif/internal/errors"

type StatusResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message,omitempty"`
	Error   *errors.ExportableError `json:"error,omitempty"`
}

func NewSuccessResponse() *StatusResponse {
	return &StatusResponse{
		Status: "success",
	}
}

func NewErrorResponse(err *errors.ExportableError) *StatusResponse {
	return &StatusResponse{
		Status: "error",
		Error:  err,
	}
}
