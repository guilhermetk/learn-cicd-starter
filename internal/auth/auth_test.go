package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey lsdkfjlkasdfklsajflaskdfjklf")

	_, err := GetAPIKey(header)

	if err != nil {
		t.Fatalf("expected no errors, got: %v", err)
	}
}
