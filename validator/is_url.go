package validator

import "net/url"

// IsURL If Not Url, then show error message
func (v *validator) IsURL(value, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	u, err := url.Parse(value)

	isURL := err == nil && u.Scheme != "" && u.Host != ""

	if !isURL {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
