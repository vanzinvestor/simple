package validator

import "net/url"

// IsURL If Not Url, then return json error
func (v *validator) IsURL(str, errorMessage string) bool {
	if _, ok := v.Errors["error"]; ok {
		return false
	}

	u, err := url.Parse(str)

	isURL := err == nil && u.Scheme != "" && u.Host != ""

	if !isURL {
		v.Errors["error"] = errMessage{message: errorMessage}.Error()
		return false
	}

	return true
}
