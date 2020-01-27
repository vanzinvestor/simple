package validator

// IsMax If value more than maximum, then show error message
func (v *validator) IsMax(value string, max int, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if value == "" {
		return true
	}

	if len(value) > max {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
