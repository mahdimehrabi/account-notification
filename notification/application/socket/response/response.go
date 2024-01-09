package response

import (
	"encoding/json"
	"io"
	"log"
)

type Status string

const (
	StatusOk            Status = "ok"
	StatusForbidden     Status = "forbidden"
	StatusInternalError Status = "internal_error"
	StatusBadRequest    Status = "bad_request"
)

type response struct {
	Data    any      `json:"data"`
	Message string   `json:"notification"`
	Errors  []string `json:"errors"`
	Status  Status   `json:"status"`
	ID      string   `json:"ID"`
}

func GenResponse(w io.Writer, data any, message string, status Status, id string, errors []string) {
	r, err := json.Marshal(response{
		Data:    data,
		Message: message,
		Errors:  errors,
		Status:  status,
		ID:      id,
	})
	if err != nil {
		log.Fatal("error")
	}

	if _, err = w.Write(r); err != nil {
		log.Fatal(err)
	}
}

func SuccessResponse(w io.Writer, data any, id string, message string) {
	GenResponse(w, data, message, StatusOk, id, []string{})
}

func InternalErrorResponse(w io.Writer, id string) {
	GenResponse(w, nil, "internal server error", StatusInternalError, id, []string{})
}

func BadRequestErrorResponse(w io.Writer, id string) {
	GenResponse(w, nil, "bad request", StatusBadRequest, id, []string{})
}
