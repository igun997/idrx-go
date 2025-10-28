package idrx

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testMethod = "POST"
)

func TestCreateSignature(t *testing.T) {
	tests := getSignatureTestCases()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSignature(tt.method, tt.url, tt.body, tt.timestamp, tt.secretKey)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type signatureTestCase struct {
	name      string
	method    string
	url       string
	body      interface{}
	timestamp string
	secretKey string
	want      string
	wantErr   bool
}

func getSignatureTestCases() []signatureTestCase {
	validCases := getValidSignatureCases()
	errorCases := getErrorSignatureCases()

	return append(validCases, errorCases...)
}

func getValidSignatureCases() []signatureTestCase {
	basicCases := getBasicValidCases()
	variationCases := getVariationCases()

	return append(basicCases, variationCases...)
}

func getBasicValidCases() []signatureTestCase {
	return []signatureTestCase{
		{
			name:      "valid signature with string body",
			method:    "POST",
			url:       "/api/v1/test",
			body:      map[string]string{"message": "hello"},
			timestamp: "1609459200",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "lD2AKBTXwjQKhGHXQfOM-gko9mnqkLqFpTlqVsVZWKU",
			wantErr:   false,
		},
		{
			name:      "valid signature with nil body",
			method:    "GET",
			url:       "/api/v1/users",
			body:      nil,
			timestamp: "1609459200",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "uGlIDt0j85RXE6YYAM7YuKFhmKm2xrFSpAjk4WWpZR4",
			wantErr:   false,
		},
		{
			name:   "valid signature with complex body",
			method: "POST",
			url:    "/api/v1/payroll",
			body: map[string]interface{}{
				"amount":   100.50,
				"currency": "USD",
				"employee": map[string]string{
					"id":   "emp123",
					"name": "John Doe",
				},
			},
			timestamp: "1609459200",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "uCj55xMxe7GpYRONMxsafLyRAY16MSFZbIlyAYdhTNI",
			wantErr:   false,
		},
	}
}

func getVariationCases() []signatureTestCase {
	return []signatureTestCase{
		{
			name:      "different methods produce different signatures",
			method:    "PUT",
			url:       "/api/v1/test",
			body:      map[string]string{"message": "hello"},
			timestamp: "1609459200",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "5lVUo4zYkBynDeuHSlXLCrPFYFdRx23X6SyRktans-0",
			wantErr:   false,
		},
		{
			name:      "different timestamps produce different signatures",
			method:    "POST",
			url:       "/api/v1/test",
			body:      map[string]string{"message": "hello"},
			timestamp: "1609459300",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "crASNC0lftOEACbZB8EJ2cnJ8C7zz-HDGovdwGYCowc",
			wantErr:   false,
		},
		{
			name:      "empty body object",
			method:    "POST",
			url:       "/api/v1/test",
			body:      map[string]interface{}{},
			timestamp: "1609459200",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "q0KzsJIceg8KNDBFJlMzNXge6NB1nye5AGIzfBxUl8I",
			wantErr:   false,
		},
	}
}

func getErrorSignatureCases() []signatureTestCase {
	return []signatureTestCase{
		{
			name:      "invalid base64 secret key",
			method:    "POST",
			url:       "/api/v1/test",
			body:      map[string]string{"message": "hello"},
			timestamp: "1609459200",
			secretKey: "invalid-base64-key!@#",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "invalid JSON body",
			method:    "POST",
			url:       "/api/v1/test",
			body:      make(chan int),
			timestamp: "1609459200",
			secretKey: "dGVzdC1zZWNyZXQta2V5LTEyMzQ1Njc4OTA=",
			want:      "",
			wantErr:   true,
		},
	}
}

func TestCreateSignatureConsistency(t *testing.T) {
	method := testMethod
	url := "/api/v1/test"
	body := map[string]string{"message": "hello"}
	timestamp := "1609459200"
	secretKey := testSecretKey()

	// Call the function multiple times with the same parameters
	sig1, err1 := CreateSignature(method, url, body, timestamp, secretKey)
	sig2, err2 := CreateSignature(method, url, body, timestamp, secretKey)
	sig3, err3 := CreateSignature(method, url, body, timestamp, secretKey)

	// All calls should succeed
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, err3)

	// All signatures should be identical
	assert.Equal(t, sig1, sig2)
	assert.Equal(t, sig2, sig3)
	assert.NotEmpty(t, sig1)
}

func TestCreateSignatureBodyOrder(t *testing.T) {
	method := testMethod
	url := "/api/v1/test"
	timestamp := "1609459200"
	secretKey := testSecretKey()

	// Test that JSON marshaling order doesn't affect signature consistency
	body1 := map[string]interface{}{
		"a": "value1",
		"b": "value2",
		"c": "value3",
	}

	body2 := map[string]interface{}{
		"c": "value3",
		"a": "value1",
		"b": "value2",
	}

	sig1, err1 := CreateSignature(method, url, body1, timestamp, secretKey)
	sig2, err2 := CreateSignature(method, url, body2, timestamp, secretKey)

	require.NoError(t, err1)
	require.NoError(t, err2)

	// Note: Go's map iteration order is not guaranteed, but JSON marshaling
	// should produce consistent results for the same key-value pairs
	assert.Equal(t, sig1, sig2)
}

func TestCreateSignatureEdgeCases(t *testing.T) {
	secretKey := testSecretKey()
	timestamp := "1609459200"

	t.Run("empty method", func(t *testing.T) {
		sig, err := CreateSignature("", "/api/v1/test", nil, timestamp, secretKey)
		require.NoError(t, err)
		assert.NotEmpty(t, sig)
	})

	t.Run("empty URL", func(t *testing.T) {
		sig, err := CreateSignature("GET", "", nil, timestamp, secretKey)
		require.NoError(t, err)
		assert.NotEmpty(t, sig)
	})

	t.Run("empty timestamp", func(t *testing.T) {
		sig, err := CreateSignature("GET", "/api/v1/test", nil, "", secretKey)
		require.NoError(t, err)
		assert.NotEmpty(t, sig)
	})

	t.Run("unicode characters in body", func(t *testing.T) {
		body := map[string]string{
			"message": "Hello 世界 🌍",
			"emoji":   "🚀💰",
		}
		sig, err := CreateSignature("POST", "/api/v1/test", body, timestamp, secretKey)
		require.NoError(t, err)
		assert.NotEmpty(t, sig)
	})
}

func BenchmarkCreateSignature(b *testing.B) {
	method := testMethod
	url := "/api/v1/payroll"
	body := map[string]interface{}{
		"amount":   100.50,
		"currency": "USD",
		"employee": map[string]string{
			"id":   "emp123",
			"name": "John Doe",
		},
	}
	timestamp := "1609459200"
	secretKey := testSecretKey()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := CreateSignature(method, url, body, timestamp, secretKey)
		if err != nil {
			b.Fatal(err)
		}
	}
}
