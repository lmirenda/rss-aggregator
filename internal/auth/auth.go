package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey Get API Key from request headers
// Authorization: ApiKey {value}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing authentication headers")
	}

	values := strings.Split(val, " ")
	if len(values) != 2 {
		return "", errors.New("malformed auth header")
	}
	if values[0] != "ApiKey" {
		return "", errors.New("malformed first part of headers")
	}
	return values[1], nil
}
