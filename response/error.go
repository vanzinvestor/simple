package response

import (
	"net/http"
)

// ERROR response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	var res map[string]string

	if err == nil {
		res = map[string]string{"error": http.StatusText(statusCode)}
		JSON(w, statusCode, res)

		return
	}

	res = map[string]string{"error": err.Error()}
	JSON(w, statusCode, res)
}
