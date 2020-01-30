package validator

// IsNotEmpty If empty, then return json error
func (v *validator) IsNotEmpty(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if str == "" {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
