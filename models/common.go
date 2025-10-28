package models

import (
	"fmt"
	"time"
)

// APIResponse represents the standard IDRX API response structure.
type APIResponse[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       T      `json:"data"`
}

// APIError represents an error returned by the IDRX API.
// It implements the error interface for consistent error handling.
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Message)
}

// ListOptions provides common pagination and filtering options.
type ListOptions struct {
	Page   int               `url:"page,omitempty"`
	Take   int               `url:"take,omitempty"`
	Sort   string            `url:"sort,omitempty"`
	Filter map[string]string `url:"-"`
}

// ListMetadata represents pagination metadata in list responses.
type ListMetadata struct {
	Page       int `json:"page"`
	Take       int `json:"take"`
	Total      int `json:"total"`
	TotalPages int `json:"totalPages"`
}

// ListResponse represents a paginated list response.
type ListResponse[T any] struct {
	Data     []T          `json:"data"`
	Metadata ListMetadata `json:"metadata"`
}

// ChainID represents supported blockchain network IDs.
type ChainID string

const (
	ChainPolygon  ChainID = "137"   // Polygon Mainnet
	ChainBSC      ChainID = "56"    // BNB Smart Chain
	ChainEthereum ChainID = "1"     // Ethereum Mainnet
	ChainArbitrum ChainID = "42161" // Arbitrum One
)

// TransactionType represents different types of transactions.
type TransactionType string

const (
	TransactionTypeMint          TransactionType = "MINT"
	TransactionTypeBurn          TransactionType = "BURN"
	TransactionTypeBridge        TransactionType = "BRIDGE"
	TransactionTypeDepositRedeem TransactionType = "DEPOSIT_REDEEM"
)

// TransactionStatus represents the status of a transaction.
type TransactionStatus string

const (
	TransactionStatusPending    TransactionStatus = "PENDING"
	TransactionStatusProcessing TransactionStatus = "PROCESSING"
	TransactionStatusCompleted  TransactionStatus = "COMPLETED"
	TransactionStatusFailed     TransactionStatus = "FAILED"
	TransactionStatusCanceled   TransactionStatus = "CANCELED"
)

// FeeType represents different types of fees.
type FeeType string

const (
	FeeTypeMint   FeeType = "MINT"
	FeeTypeRedeem FeeType = "REDEEM"
	FeeTypeRefund FeeType = "REFUND"
)

// RequestType represents different request types for minting.
type RequestType string

const (
	RequestTypeIDRX RequestType = "idrx"
	RequestTypeUSDT RequestType = "usdt"
)

// Fee represents a fee structure.
type Fee struct {
	Name     string `json:"name"`
	Amount   string `json:"amount"`
	IsActive bool   `json:"isActive"`
}

// DeleteResponse represents a standard delete operation response.
type DeleteResponse struct {
	Data any `json:"data"` // Usually null for delete operations
}

// TimestampedData represents data with creation/update timestamps.
type TimestampedData struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
