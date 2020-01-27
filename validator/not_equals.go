package validator

// NotEquals If Equal to, then show error message
func (v *validator) NotEquals(value, match, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if value == match {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
