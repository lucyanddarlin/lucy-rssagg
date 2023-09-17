package auth

import (
	"errors"
	"net/http"
	"strings"
)

const HEADER_KEY = "APIKey"

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authorization info is found")
	}

	nextVal := strings.Split(val, " ")
	if len(nextVal) != 2 {
		return "", errors.New("malformed auth header")
	}
	if nextVal[0] != HEADER_KEY {
		return "", errors.New("malformed first part of auth header")
	}
	return nextVal[1], nil
}
