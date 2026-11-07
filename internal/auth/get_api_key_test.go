package auth

import (
	"log"
	"net/http"
	"testing"
)

func TestApiKeyFromAuth(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey Test1234")

	got, err := GetAPIKey(headers)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	want := "Test1234"

	if got != want {
		t.Errorf("GetApiKey() = %q, want %q", got, want)
	}
}
