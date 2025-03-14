package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ConvoyResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func DispatchSuccess(w http.ResponseWriter, message string, data any) {
	response := ConvoyResponse{
		Status:  "successful",
		Message: message,
		Data:    data,
	}

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Fprint(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func DispatchError(w http.ResponseWriter, message string) {
	response := ConvoyResponse{
		Status:  "failed",
		Message: message,
		Data:    nil,
	}

	res, err := json.Marshal(response)
	if err != nil {
		fmt.Fprint(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write(res)
	return
}
