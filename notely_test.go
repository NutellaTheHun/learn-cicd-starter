package main

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestAuth(t *testing.T) {
	basicH := createTestHeader("Authorization", "ApiKey 69420")
	noValH := createTestHeader("Authorization", "")
	badValSizeH := createTestHeader("Authorization", "ApiKey")
	tests := map[string]struct {
		input http.Header
		want  error
	}{
		"basic":        {input: basicH, want: nil},
		"no value":     {input: noValH, want: errors.New("no authorization header included")},
		"bad val size": {input: badValSizeH, want: errors.New("malformed authorization header")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := auth.GetAPIKey(tc.input)
			if err != nil {
				if !reflect.DeepEqual(tc.want, err) {
					t.Fatalf("exected: %#v, got: %#v", tc.want, err)
				}
			}
			if !reflect.DeepEqual(tc.want, err) {
				t.Fatalf("exected: %#v, got: %#v", tc.want, err)
			}
		})
	}
}

func createTestHeader(headerKey, val string) http.Header {
	result := http.Header{}
	result.Set(headerKey, val)
	return result
}
