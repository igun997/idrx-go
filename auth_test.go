package idrx

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

// testSecretKey returns a valid base64-encoded secret key for testing
func testSecretKey() string {
	return base64.StdEncoding.EncodeToString([]byte("test-secret-key"))
}

func TestNewBusinessAuth(t *testing.T) {
	apiKey := "test-business-api-key"
	secretKey := testSecretKey()

	auth := NewBusinessAuth(apiKey, secretKey)
	businessAuth, ok := auth.(*BusinessAuth)

	if !ok {
		t.Fatal("expected *BusinessAuth type")
	}

	if businessAuth.apiKey != apiKey {
		t.Errorf("expected apiKey %s, got %s", apiKey, businessAuth.apiKey)
	}

	if businessAuth.secretKey != secretKey {
		t.Errorf("expected secretKey %s, got %s", secretKey, businessAuth.secretKey)
	}

	if auth.GetAuthType() != "business" {
		t.Errorf("expected auth type 'business', got %s", auth.GetAuthType())
	}
}

func TestNewUserAuth(t *testing.T) {
	apiKey := "test-user-api-key"
	secretKey := testSecretKey()

	auth := NewUserAuth(apiKey, secretKey)
	userAuth, ok := auth.(*UserAuth)

	if !ok {
		t.Fatal("expected *UserAuth type")
	}

	if userAuth.apiKey != apiKey {
		t.Errorf("expected apiKey %s, got %s", apiKey, userAuth.apiKey)
	}

	if userAuth.secretKey != secretKey {
		t.Errorf("expected secretKey %s, got %s", secretKey, userAuth.secretKey)
	}

	if auth.GetAuthType() != "user" {
		t.Errorf("expected auth type 'user', got %s", auth.GetAuthType())
	}
}

func TestBusinessAuthAuthenticate(t *testing.T) {
	auth := NewBusinessAuth("test-api-key", testSecretKey())

	req, err := http.NewRequest("GET", "https://idrx.co/api/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	err = auth.Authenticate(ctx, req, nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify headers are set
	if req.Header.Get("idrx-api-key") != "test-api-key" {
		t.Errorf("expected idrx-api-key header to be set")
	}

	if req.Header.Get("idrx-api-sig") == "" {
		t.Errorf("expected idrx-api-sig header to be set")
	}

	if req.Header.Get("idrx-api-ts") == "" {
		t.Errorf("expected idrx-api-ts header to be set")
	}
}

func TestUserAuthAuthenticate(t *testing.T) {
	auth := NewUserAuth("test-user-api-key", testSecretKey())

	req, err := http.NewRequest("POST", "https://idrx.co/api/v1/user/test", strings.NewReader(`{"test": "data"}`))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	err = auth.Authenticate(ctx, req, map[string]string{"test": "data"})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Verify headers are set
	if req.Header.Get("idrx-api-key") != "test-user-api-key" {
		t.Errorf("expected idrx-api-key header to be set")
	}

	if req.Header.Get("idrx-api-sig") == "" {
		t.Errorf("expected idrx-api-sig header to be set")
	}

	if req.Header.Get("idrx-api-ts") == "" {
		t.Errorf("expected idrx-api-ts header to be set")
	}
}

func TestAuthenticateWithCanceledContext(t *testing.T) {
	auth := NewBusinessAuth("test-api-key", testSecretKey())

	req, err := http.NewRequest("GET", "https://idrx.co/api/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err = auth.Authenticate(ctx, req, nil)

	if err == nil {
		t.Error("expected error with canceled context")
	}

	if !strings.Contains(err.Error(), "context canceled") {
		t.Errorf("expected context cancellation error, got %v", err)
	}
}

func TestAuthenticateWithTimeout(t *testing.T) {
	auth := NewBusinessAuth("test-api-key", testSecretKey())

	req, err := http.NewRequest("GET", "https://idrx.co/api/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()

	// Wait for context to timeout
	time.Sleep(1 * time.Millisecond)

	err = auth.Authenticate(ctx, req, nil)

	if err == nil {
		t.Error("expected error with timed out context")
	}
}

func TestNewNoAuth(t *testing.T) {
	auth := NewNoAuth()
	noAuth, ok := auth.(*NoAuth)

	if !ok {
		t.Fatal("expected *NoAuth type")
	}

	if auth.GetAuthType() != "none" {
		t.Errorf("expected auth type 'none', got %s", auth.GetAuthType())
	}

	// Test that authenticate does nothing
	req, err := http.NewRequest("GET", "https://idrx.co/api/v1/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	err = noAuth.Authenticate(ctx, req, nil)
	if err != nil {
		t.Errorf("expected no error from NoAuth, got %v", err)
	}

	// Verify no headers are set
	if req.Header.Get("idrx-api-key") != "" {
		t.Error("NoAuth should not set any headers")
	}
}

func TestAuthProviderInterface(t *testing.T) {
	// Test that all auth types implement the interface
	var providers []AuthProvider = []AuthProvider{
		NewBusinessAuth("key", testSecretKey()),
		NewUserAuth("key", testSecretKey()),
		NewNoAuth(),
	}

	expectedTypes := []string{"business", "user", "none"}

	for i, provider := range providers {
		if provider.GetAuthType() != expectedTypes[i] {
			t.Errorf("expected type %s, got %s", expectedTypes[i], provider.GetAuthType())
		}

		// Test that authenticate method exists and can be called
		req, _ := http.NewRequest("GET", "https://test.com", nil)
		provider.Authenticate(context.Background(), req, nil)
	}
}

func TestBusinessAuthHeaderGeneration(t *testing.T) {
	auth := NewBusinessAuth("test-key", testSecretKey())

	testURL, _ := url.Parse("https://idrx.co/api/v1/members")
	req := &http.Request{
		Method: "GET",
		URL:    testURL,
		Header: make(http.Header),
	}

	err := auth.Authenticate(context.Background(), req, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Verify all required headers are present
	requiredHeaders := []string{"idrx-api-key", "idrx-api-sig", "idrx-api-ts"}
	for _, header := range requiredHeaders {
		if req.Header.Get(header) == "" {
			t.Errorf("missing required header: %s", header)
		}
	}

	// Verify API key is correct
	if req.Header.Get("idrx-api-key") != "test-key" {
		t.Errorf("incorrect API key in header")
	}
}

func TestUserAuthHeaderGeneration(t *testing.T) {
	auth := NewUserAuth("user-key", testSecretKey())

	testURL, _ := url.Parse("https://idrx.co/api/v1/user/profile")
	req := &http.Request{
		Method: "POST",
		URL:    testURL,
		Header: make(http.Header),
	}

	body := map[string]string{"field": "value"}
	err := auth.Authenticate(context.Background(), req, body)
	if err != nil {
		t.Fatal(err)
	}

	// Verify all required headers are present
	requiredHeaders := []string{"idrx-api-key", "idrx-api-sig", "idrx-api-ts"}
	for _, header := range requiredHeaders {
		if req.Header.Get(header) == "" {
			t.Errorf("missing required header: %s", header)
		}
	}

	// Verify API key is correct
	if req.Header.Get("idrx-api-key") != "user-key" {
		t.Errorf("incorrect API key in header")
	}
}
