package idrx

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// AuthProvider defines the interface for authentication strategies.
// This follows the strategy pattern to support different auth mechanisms.
type AuthProvider interface {
	// Authenticate adds authentication headers to the HTTP request
	Authenticate(ctx context.Context, req *http.Request, body any) error

	// GetAuthType returns the type of authentication (business, user, etc.)
	GetAuthType() string
}

// BusinessAuth provides authentication for business-level API operations.
// This is used for general endpoints and organization management.
type BusinessAuth struct {
	apiKey    string
	secretKey string
}

// NewBusinessAuth creates a new business-level authentication provider.
func NewBusinessAuth(apiKey, secretKey string) AuthProvider {
	return &BusinessAuth{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// Authenticate implements the AuthProvider interface for business auth.
func (b *BusinessAuth) Authenticate(ctx context.Context, req *http.Request, body any) error {
	// Check if context is already canceled
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled during authentication: %w", err)
	}
	return b.addAuthHeaders(ctx, req, body)
}

// GetAuthType returns the authentication type.
func (b *BusinessAuth) GetAuthType() string {
	return "business"
}

// addAuthHeaders adds the required authentication headers for IDRX API.
func (b *BusinessAuth) addAuthHeaders(ctx context.Context, req *http.Request, body any) error {
	// Check context before proceeding with potentially time-consuming operations
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled before signature generation: %w", err)
	}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// Generate signature using existing CreateSignature function
	signature, err := CreateSignature(req.Method, req.URL.Path, body, timestamp, b.secretKey)
	if err != nil {
		return fmt.Errorf("failed to create signature: %w", err)
	}

	// Check context again before setting headers
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled during header setup: %w", err)
	}

	// Add required IDRX authentication headers
	req.Header.Set("idrx-api-key", b.apiKey)
	req.Header.Set("idrx-api-sig", signature)
	req.Header.Set("idrx-api-ts", timestamp)

	return nil
}

// UserAuth provides authentication for user-level API operations.
// This is used for user-specific operations and is obtained from onboarding response.
type UserAuth struct {
	apiKey    string
	secretKey string
}

// NewUserAuth creates a new user-level authentication provider.
func NewUserAuth(apiKey, secretKey string) AuthProvider {
	return &UserAuth{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// Authenticate implements the AuthProvider interface for user auth.
func (u *UserAuth) Authenticate(ctx context.Context, req *http.Request, body any) error {
	// Check if context is already canceled
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled during authentication: %w", err)
	}
	return u.addAuthHeaders(ctx, req, body)
}

// GetAuthType returns the authentication type.
func (u *UserAuth) GetAuthType() string {
	return "user"
}

// addAuthHeaders adds the required authentication headers for IDRX API.
func (u *UserAuth) addAuthHeaders(ctx context.Context, req *http.Request, body any) error {
	// Check context before proceeding with potentially time-consuming operations
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled before signature generation: %w", err)
	}
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// Generate signature using existing CreateSignature function
	signature, err := CreateSignature(req.Method, req.URL.Path, body, timestamp, u.secretKey)
	if err != nil {
		return fmt.Errorf("failed to create signature: %w", err)
	}

	// Check context again before setting headers
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context canceled during header setup: %w", err)
	}

	// Add required IDRX authentication headers
	req.Header.Set("idrx-api-key", u.apiKey)
	req.Header.Set("idrx-api-sig", signature)
	req.Header.Set("idrx-api-ts", timestamp)

	return nil
}

// NoAuth provides a no-op authentication provider for testing or public endpoints.
type NoAuth struct{}

// NewNoAuth creates a new no-authentication provider.
func NewNoAuth() AuthProvider {
	return &NoAuth{}
}

// Authenticate implements the AuthProvider interface with no-op.
func (n *NoAuth) Authenticate(_ context.Context, _ *http.Request, _ any) error {
	return nil
}

// GetAuthType returns the authentication type.
func (n *NoAuth) GetAuthType() string {
	return "none"
}
