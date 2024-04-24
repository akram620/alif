package errors

import (
	"net/http"
	"strings"
)

const (
	EXPECTED   = "expected"
	UNEXPECTED = "unexpected"
	WARNING    = "warning"
	FATAL      = "fatal"
)

type ExportableError struct {
	Code            string `json:"error"`
	CodeDescription string `json:"error_description"`
	HttpStatus      int    `json:"http_status"`
	Severity        string `json:"severity"`
	Message         string `json:"message"`
	Caller          string `json:"caller"`
}

func (e *ExportableError) WithMessage(message string) *ExportableError {
	skip := 2
	caller := GetCaller(skip)

	for strings.Contains(caller, "errors.go") {
		skip += 1
		caller = GetCaller(skip)
	}

	errorWithMessage := *e
	errorWithMessage.Message = message
	errorWithMessage.Caller = caller

	return &errorWithMessage
}

func (e *ExportableError) Default() *ExportableError {
	return e.WithMessage(e.Message)
}

var ErrBadRequestParseBody = ExportableError{
	Code:            "4006",
	CodeDescription: "ERROR_PARSE_BODY",
	HttpStatus:      http.StatusBadRequest,
	Severity:        EXPECTED,
}

var ErrInternalServerError = ExportableError{
	Code:            "5000",
	CodeDescription: "ERROR_INTERNAL_SERVER_ERROR",
	HttpStatus:      http.StatusInternalServerError,
	Severity:        FATAL,
}

var ErrInternalServerErrorDatabaseFailed = ExportableError{
	Code:            "5003",
	CodeDescription: "ERROR_INTERNAL_SERVER_ERROR_DATABASE_OPERATION_FAILED",
	HttpStatus:      http.StatusInternalServerError,
	Severity:        FATAL,
}
