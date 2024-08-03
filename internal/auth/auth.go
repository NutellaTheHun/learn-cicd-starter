package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

/*
func TestAuth(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"basic":        {input: createTestHeader("Authorization", "ApiKey  69420"), want: "69420"},
		"no value":     {input: createTestHeader("Authorization", ""), want: "no authorization header included"},
		"bad val size": {input: createTestHeader("Authorization", "ApiKey"), want: "malformed authorization header"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("exected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func createTestHeader(headerKey, val string) http.Header {
	result := http.Header{}
	result.Set(headerKey, val)
	return result
}
*/
