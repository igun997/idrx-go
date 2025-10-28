package models

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// BlockchainTransaction represents a blockchain transaction with IDRX-specific metadata
type BlockchainTransaction struct {
	Hash        common.Hash    `json:"hash"`
	ChainID     uint64         `json:"chainId"`
	From        common.Address `json:"from"`
	To          common.Address `json:"to"`
	Value       *big.Int       `json:"value"`
	Gas         uint64         `json:"gas"`
	GasPrice    *big.Int       `json:"gasPrice"`
	GasUsed     uint64         `json:"gasUsed,omitempty"`
	Status      uint64         `json:"status,omitempty"`
	BlockNumber uint64         `json:"blockNumber,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
}

// BridgeTransaction represents a cross-chain bridge transaction
type BridgeTransaction struct {
	ID              string          `json:"id"`
	FromChainID     uint64          `json:"fromChainId"`
	ToChainID       uint64          `json:"toChainId"`
	FromAddress     common.Address  `json:"fromAddress"`
	ToAddress       common.Address  `json:"toAddress"`
	Amount          decimal.Decimal `json:"amount"`
	AmountAfterFees decimal.Decimal `json:"amountAfterFees"`
	PlatformFee     decimal.Decimal `json:"platformFee"`
	BridgeNonce     *big.Int        `json:"bridgeNonce"`
	BurnTxHash      common.Hash     `json:"burnTxHash,omitempty"`
	MintTxHash      common.Hash     `json:"mintTxHash,omitempty"`
	Status          BridgeStatus    `json:"status"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
	CompletedAt     *time.Time      `json:"completedAt,omitempty"`
}

// BridgeStatus represents the status of a bridge transaction
type BridgeStatus string

const (
	BridgeStatusRequested BridgeStatus = "REQUESTED"
	BridgeStatusBurning   BridgeStatus = "BURNING"
	BridgeStatusBurned    BridgeStatus = "BURNED"
	BridgeStatusMinting   BridgeStatus = "MINTING"
	BridgeStatusCompleted BridgeStatus = "COMPLETED"
	BridgeStatusFailed    BridgeStatus = "FAILED"
	BridgeStatusCancelled BridgeStatus = "CANCELED"
)

// BlockchainMintRequest represents a token minting request with blockchain-specific data
type BlockchainMintRequest struct {
	ID               string          `json:"id"`
	ChainID          uint64          `json:"chainId"`
	ToAddress        common.Address  `json:"toAddress"`
	Amount           decimal.Decimal `json:"amount"`
	TxHash           common.Hash     `json:"txHash,omitempty"`
	Status           MintStatus      `json:"status"`
	RequestType      string          `json:"requestType"` // "idrx", "usdt", "usdc", etc.
	PaymentReference string          `json:"paymentReference,omitempty"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
	CompletedAt      *time.Time      `json:"completedAt,omitempty"`
}

// MintStatus represents the status of a mint request
type MintStatus string

const (
	MintStatusRequested       MintStatus = "REQUESTED"
	MintStatusPaymentPending  MintStatus = "PAYMENT_PENDING"
	MintStatusPaymentReceived MintStatus = "PAYMENT_RECEIVED"
	MintStatusMinting         MintStatus = "MINTING"
	MintStatusCompleted       MintStatus = "COMPLETED"
	MintStatusFailed          MintStatus = "FAILED"
	MintStatusExpired         MintStatus = "EXPIRED"
	MintStatusCancelled       MintStatus = "CANCELED"
)

// BlockchainRedeemRequest represents a token redemption request with blockchain-specific data
type BlockchainRedeemRequest struct {
	ID                string          `json:"id"`
	ChainID           uint64          `json:"chainId"`
	FromAddress       common.Address  `json:"fromAddress"`
	Amount            decimal.Decimal `json:"amount"`
	BankAccountNumber string          `json:"bankAccountNumber"`
	BankCode          string          `json:"bankCode"`
	BankName          string          `json:"bankName"`
	BankAccountName   string          `json:"bankAccountName"`
	TxHash            common.Hash     `json:"txHash,omitempty"`
	Status            RedeemStatus    `json:"status"`
	CustRefNumber     string          `json:"custRefNumber,omitempty"`
	DisburseID        *big.Int        `json:"disburseId,omitempty"`
	Notes             string          `json:"notes,omitempty"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
	CompletedAt       *time.Time      `json:"completedAt,omitempty"`
}

// RedeemStatus represents the status of a redeem request
type RedeemStatus string

const (
	RedeemStatusRequested  RedeemStatus = "REQUESTED"
	RedeemStatusBurning    RedeemStatus = "BURNING"
	RedeemStatusBurned     RedeemStatus = "BURNED"
	RedeemStatusProcessing RedeemStatus = "PROCESSING"
	RedeemStatusCompleted  RedeemStatus = "COMPLETED"
	RedeemStatusFailed     RedeemStatus = "FAILED"
	RedeemStatusCancelled  RedeemStatus = "CANCELED"
)

// TokenBalance represents a token balance for a specific address and chain
type TokenBalance struct {
	ChainID     uint64          `json:"chainId"`
	Address     common.Address  `json:"address"`
	Balance     decimal.Decimal `json:"balance"`
	LastUpdated time.Time       `json:"lastUpdated"`
}

// NetworkStatus represents the status of a blockchain network
type NetworkStatus struct {
	ChainID      uint64    `json:"chainId"`
	Name         string    `json:"name"`
	IsHealthy    bool      `json:"isHealthy"`
	BlockNumber  uint64    `json:"blockNumber"`
	LastChecked  time.Time `json:"lastChecked"`
	ResponseTime int64     `json:"responseTime"` // in milliseconds
	ErrorMessage string    `json:"errorMessage,omitempty"`
}

// EventLog represents a blockchain event log
type EventLog struct {
	ID          string         `json:"id"`
	ChainID     uint64         `json:"chainId"`
	Address     common.Address `json:"address"`
	Topics      []common.Hash  `json:"topics"`
	Data        []byte         `json:"data"`
	BlockNumber uint64         `json:"blockNumber"`
	TxHash      common.Hash    `json:"txHash"`
	TxIndex     uint           `json:"txIndex"`
	LogIndex    uint           `json:"logIndex"`
	EventName   string         `json:"eventName"`
	ParsedData  interface{}    `json:"parsedData,omitempty"`
	CreatedAt   time.Time      `json:"createdAt"`
}

// BurnBridgeEvent represents a parsed BurnBridge event
type BurnBridgeEvent struct {
	User           common.Address  `json:"user"`
	Amount         decimal.Decimal `json:"amount"`
	AmountAfterCut decimal.Decimal `json:"amountAfterCut"`
	ToChain        uint64          `json:"toChain"`
	BridgeNonce    *big.Int        `json:"bridgeNonce"`
	PlatformFee    decimal.Decimal `json:"platformFee"`
	BlockNumber    uint64          `json:"blockNumber"`
	TxHash         common.Hash     `json:"txHash"`
}

// MintBridgeEvent represents a parsed MintBridge event
type MintBridgeEvent struct {
	User            common.Address  `json:"user"`
	Amount          decimal.Decimal `json:"amount"`
	AmountAfterCut  decimal.Decimal `json:"amountAfterCut"`
	FromChain       uint64          `json:"fromChain"`
	FromBridgeNonce *big.Int        `json:"fromBridgeNonce"`
	PlatformFee     decimal.Decimal `json:"platformFee"`
	BlockNumber     uint64          `json:"blockNumber"`
	TxHash          common.Hash     `json:"txHash"`
}

// BurnWithAccountNumberEvent represents a parsed BurnWithAccountNumber event
type BurnWithAccountNumberEvent struct {
	User                common.Address  `json:"user"`
	Amount              decimal.Decimal `json:"amount"`
	HashedAccountNumber string          `json:"hashedAccountNumber"`
	BlockNumber         uint64          `json:"blockNumber"`
	TxHash              common.Hash     `json:"txHash"`
}

// TransferEvent represents a parsed Transfer event
type TransferEvent struct {
	From        common.Address  `json:"from"`
	To          common.Address  `json:"to"`
	Value       decimal.Decimal `json:"value"`
	BlockNumber uint64          `json:"blockNumber"`
	TxHash      common.Hash     `json:"txHash"`
}

// WebhookPayload represents a webhook payload for blockchain events
type WebhookPayload struct {
	Type      string      `json:"type"`
	ChainID   uint64      `json:"chainId"`
	Event     string      `json:"event"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
	Signature string      `json:"signature,omitempty"`
}

// GasEstimate represents gas estimation data
type GasEstimate struct {
	GasLimit    uint64    `json:"gasLimit"`
	GasPrice    *big.Int  `json:"gasPrice"`
	TotalCost   *big.Int  `json:"totalCost"`
	EstimatedAt time.Time `json:"estimatedAt"`
}

// ContractCallResult represents the result of a contract call
type ContractCallResult struct {
	Success    bool        `json:"success"`
	ReturnData interface{} `json:"returnData,omitempty"`
	GasUsed    uint64      `json:"gasUsed,omitempty"`
	Error      string      `json:"error,omitempty"`
	TxHash     common.Hash `json:"txHash,omitempty"`
}

// ChainStats represents statistics for a blockchain network
type ChainStats struct {
	ChainID           uint64          `json:"chainId"`
	TotalSupply       decimal.Decimal `json:"totalSupply"`
	TotalHolders      uint64          `json:"totalHolders"`
	TotalTransactions uint64          `json:"totalTransactions"`
	LastBlockNumber   uint64          `json:"lastBlockNumber"`
	AvgBlockTime      time.Duration   `json:"avgBlockTime"`
	UpdatedAt         time.Time       `json:"updatedAt"`
}
