// Package util utilities including auth, reponse, error handling
package util

import (
	"encoding/json"
	"net/http"
)

// Payload payload type
type Payload interface{}

// RespondWithError response with error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON response with json
func RespondWithJSON(w http.ResponseWriter, code int, payload Payload) {
	w.WriteHeader(code)
	b, err := json.Marshal(payload)
	if err != nil {
		RespondWithError(w, code, err.Error())
	}
	w.Write(b)
	w.Header().Set("Content-Type", "application/json")
}
