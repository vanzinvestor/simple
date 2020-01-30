package validator

// Exists If empty, then return json error
func (v *validator) Exists(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if str == "" {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
