package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JsonResponse returns a JSON response representation
func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		return
	}

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

// ErrorResponse returns a json error representation
func ErrorResponse(w http.ResponseWriter, statusCode int, error error) {
	JsonResponse(w, statusCode, struct {
		Error string `json:"error"`
	}{Error: error.Error()})
}
