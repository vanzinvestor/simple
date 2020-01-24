package validator

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vanzinvestor/simple/response"
)

type validator struct {
	Errors map[string]string
}

// PayloadValidation payload validation
type PayloadValidation interface {
	IsValid() (bool, map[string]string)
}

// NewValidator use valid field
func NewValidator() *validator {
	return &validator{Errors: make(map[string]string)}
}

func (v *validator) IsValid() bool {
	return len(v.Errors) == 0
}

// ValidatePayload validate payload
func ValidatePayload(next httprouter.Handle, payload PayloadValidation) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			response.ERROR(w, http.StatusBadRequest, err)

			return
		}

		defer r.Body.Close()

		if isValid, errs := payload.IsValid(); !isValid {
			response.JSON(w, http.StatusBadRequest, errs)

			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)
		next(w, r.WithContext(ctx), ps)
	}
}
