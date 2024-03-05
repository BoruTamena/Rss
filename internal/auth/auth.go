package auth

import (
	"errors"
	"net/http"
	"strings"
)

// getapikey form
// header of http request
// example: ApiKey {apikey}
func GetApiKey(header http.Header) (string, error) {

	val := header.Get("authorization")

	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] == "ApiKey" {
		return "", errors.New("malformed frist part of auth header ")
	}

	return vals[1], nil
}
