package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type responseType struct {
	Status string         `json:"status"`
	Data   any            `json:"data,omitempty"`
	Error  *errorResponse `json:"error,omitempty"`
}

func ParseJSON(w http.ResponseWriter, r *http.Request, payload any) error {
	if r.ContentLength == 0 {
		return fmt.Errorf("request body not present")
	}

	var err error
	defer func() {
		if closeErr := r.Body.Close(); closeErr != nil {
			if err == nil {
				err = fmt.Errorf("error closing response body: %w", closeErr)
			} else {
				err = fmt.Errorf("%w; error closing response body: %v", err, closeErr)
			}
		}
	}()

	maxBytesReader := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytesReader))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(payload)
}

func writeJSON(w http.ResponseWriter, statusCode int, payload responseType) {
	w.Header().Set("Content-Type", "applicaton/json")
	w.WriteHeader(statusCode)
	enc := json.NewEncoder(w)
	if err := enc.Encode(payload); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, payload any) {
	writeJSON(w, statusCode, responseType{
		Status: "success",
		Data:   payload,
	})
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, responseType{
		Status: "error",
		Error: &errorResponse{
			Code: statusCode,
			Msg:  message,
		},
	})
}
