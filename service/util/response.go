// Package util utilities including auth, reponse, error handling
package util

import (
	"bytes"
	"encoding/json"
	"net/http"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Payload payload type
type Payload interface{}

// RespondWithError response with error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON response with json
func RespondWithJSON(w http.ResponseWriter, code int, payload Payload) {
	buf := new(bytes.Buffer)
	newwriter := transform.NewWriter(buf, japanese.ShiftJIS.NewEncoder())

	defer buf.Reset()
	defer newwriter.Close()

	w.WriteHeader(code)
	json.NewEncoder(newwriter).Encode(payload)
	w.Write(buf.Bytes())
	w.Header().Set("Content-Type", "application/json")
}
