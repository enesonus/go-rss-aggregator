package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts the API key from the request headers
// Example:
// Authorization: ApiKey (APIKEY HERE)
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No Authorization value found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Invalid Authorization value")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Invalid First Authorization value")
	}
	return vals[1], nil
}