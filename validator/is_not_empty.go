package validator

// IsNotEmpty If empty, then show error message
func (v *validator) IsNotEmpty(value, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if value == "" {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
