// Package idrx provides integration to IDRX.co REST API.
package idrx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// CreateSignature create a signature for the given method, URL, body, timestamp, and secret key
// for API Call to IDRX
func CreateSignature(method, url string, body interface{}, timestamp, secretKey string) (string, error) {
	bodyBuffer, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	secret, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(timestamp))
	h.Write([]byte(method))
	h.Write([]byte(url))

	if bodyBuffer != nil {
		h.Write(bodyBuffer)
	}

	hash := h.Sum(nil)
	signature := base64.RawURLEncoding.EncodeToString(hash)

	return signature, nil
}
