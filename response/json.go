package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if data == nil {
		data = map[string]string{}
	}

	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	w.Write(buf.Bytes())
	return
}
