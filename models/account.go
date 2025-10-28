package models

import (
	"os"
	"time"
)

// OnboardingRequest represents the request to onboard a new user.
type OnboardingRequest struct {
	Email    string   `json:"email" validate:"required,email"`
	Fullname string   `json:"fullname" validate:"required"`
	Address  string   `json:"address" validate:"required"`
	IDNumber string   `json:"idNumber" validate:"required"`
	IDFile   *os.File `json:"-"` // File is handled separately in multipart form
}

// OnboardingResponse represents the response from user onboarding.
type OnboardingResponse struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	APIKey    string    `json:"apiKey"`
	APISecret string    `json:"apiSecret"`
}

// Member represents a member in the organization.
type Member struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	CreatedAt time.Time `json:"createdAt"`
	APIKeys   []APIKey  `json:"ApiKeys,omitempty"`
}

// APIKey represents an API key associated with a member.
type APIKey struct {
	ID        int       `json:"id"`
	Key       string    `json:"key"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

// MembersResponse represents the response from getting organization members.
type MembersResponse = APIResponse[[]Member]

// AddBankAccountRequest represents the request to add a bank account.
type AddBankAccountRequest struct {
	BankAccountNumber string `json:"bankAccountNumber" validate:"required"`
	BankCode          string `json:"bankCode" validate:"required"`
}

// BankAccount represents a bank account with deposit information.
type BankAccount struct {
	ID                   int           `json:"id"`
	BankAccountNumber    string        `json:"bankAccountNumber"`
	BankCode             string        `json:"bankCode"`
	BankName             string        `json:"bankName"`
	IsActive             bool          `json:"isActive"`
	CreatedAt            time.Time     `json:"createdAt"`
	UpdatedAt            time.Time     `json:"updatedAt"`
	DepositWalletAddress DepositWallet `json:"DepositWalletAddress"`
}

// DepositWallet represents deposit wallet information for a bank account.
type DepositWallet struct {
	ID           int       `json:"id"`
	Address      string    `json:"address"`
	ChainID      string    `json:"chainId"`
	ChainName    string    `json:"chainName"`
	TokenAddress string    `json:"tokenAddress"`
	TokenSymbol  string    `json:"tokenSymbol"`
	IsActive     bool      `json:"isActive"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// BankAccountResponse represents the response from adding a bank account.
type BankAccountResponse = APIResponse[BankAccount]

// BankAccountsResponse represents the response from getting bank accounts.
type BankAccountsResponse = APIResponse[[]BankAccount]

// Bank represents a supported bank with transfer limits.
type Bank struct {
	BankCode          string `json:"bankCode"`
	BankName          string `json:"bankName"`
	MaxAmountTransfer string `json:"maxAmountTransfer"`
}

// BanksResponse represents the response from getting supported banks.
type BanksResponse = APIResponse[[]Bank]
