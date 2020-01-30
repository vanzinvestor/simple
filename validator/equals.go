package validator

// Equals If Not Equal, then return json error
func (v *validator) Equals(str1, str2, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	if str1 != str2 {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
