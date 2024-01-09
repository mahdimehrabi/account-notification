package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Data    any      `json:"data"`
	Message string   `json:"message"`
	Status  int      `json:"status"`
	Errors  []string `json:"errors"`
}

func GenResponse(w http.ResponseWriter, status int, data any, message string, errors []string) {
	r, err := json.Marshal(response{
		Data:    data,
		Message: message,
		Status:  status,
		Errors:  errors,
	})
	if err != nil {
		log.Fatal("error")
	}

	w.WriteHeader(status)
	if _, err = w.Write(r); err != nil {
		log.Fatal(err)
	}
}

func SuccessResponse(w http.ResponseWriter, data any, message string) {
	GenResponse(w, http.StatusOK, data, message, []string{})
}

func InternalErrorResponse(w http.ResponseWriter) {
	GenResponse(w, http.StatusInternalServerError, nil, "internal server error", []string{})
}

func BadRequestErrorResponse(w http.ResponseWriter) {
	GenResponse(w, http.StatusBadRequest, nil, "bad request", []string{})
}

func MethodNotAllowedErrorResponse(w http.ResponseWriter) {
	GenResponse(w, http.StatusMethodNotAllowed, nil, "method not allowed", []string{})
}
