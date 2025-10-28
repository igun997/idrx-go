# IDRX Smart Contract Integration

This package provides comprehensive smart contract integration for the IDRX stablecoin ecosystem, enabling direct blockchain operations alongside the existing REST API functionality.

## Features

- **Multi-Chain Support**: Base, Polygon, BNB Smart Chain, Lisk, Kaia, World Chain, Etherlink, and Gnosis
- **Complete Contract Operations**: Transfer, mint, burn, bridge, and redemption
- **Event Monitoring**: Real-time blockchain event tracking
- **Gas Optimization**: Intelligent gas estimation and price management
- **Failover Support**: Multiple RPC endpoints per network for reliability
- **Type Safety**: Full Go type safety with generated contract bindings

## Quick Start

### 1. Basic Setup

```go
import "github.com/widnyana/idrx-go"

// Create client with blockchain capabilities
client := idrx.NewClient(
    idrx.WithUserAuth(apiKey, secretKey),
    idrx.WithBlockchain(privateKeyHex), // Enables blockchain operations
)
```

### 2. Check Token Balance

```go
// Using network names (recommended)
balance, err := client.Blockchain.GetBalanceByNetwork(ctx, blockchain.BaseMainnet, walletAddress)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Balance: %s IDRX", balance.String())

// Or using chain ID (legacy)
balance, err := client.Blockchain.GetBalance(ctx, 8453, walletAddress)
```

### 3. Transfer Tokens

```go
result, err := client.Blockchain.Transfer(ctx, 8453, recipientAddress, "1000.50")
if err != nil {
    log.Fatal(err)
}

// Wait for confirmation
status, err := client.Blockchain.WaitForTransaction(ctx, 8453, result.TxHash)
```

### 4. Cross-Chain Bridge

```go
bridgeRequest := &idrx.BridgeTransactionRequest{
    FromChainID: 8453,  // Base
    ToChainID:   137,   // Polygon
    ToAddress:   recipientAddress,
    Amount:      "5000.0",
}

result, err := client.Blockchain.InitiateBridge(ctx, bridgeRequest)
```

### 5. Burn for Redemption

```go
result, err := client.Blockchain.BurnForRedemption(ctx, 8453, "1000.0", bankAccountNumber)
```

## Supported Networks

| Network Identifier  | Network         | Chain ID | Contract Address                             |
| ------------------- | --------------- | -------- | -------------------------------------------- |
| `BaseMainnet`       | Base            | 8453     | `0x18Bc5bcC660cf2B9cE3cd51a404aFe1a0cBD3C22` |
| `PolygonMainnet`    | Polygon         | 137      | `0x649a2DA7B28E0D54c13D5eFf95d3A660652742cC` |
| `BSCMainnet`        | BNB Smart Chain | 56       | `0x649a2DA7B28E0D54c13D5eFf95d3A660652742cC` |
| `LiskMainnet`       | Lisk            | 1135     | `0x18Bc5bcC660cf2B9cE3cd51a404aFe1a0cBD3C22` |
| `KaiaMainnet`       | Kaia            | 8217     | `0x18Bc5bcC660cf2B9cE3cd51a404aFe1a0cBD3C22` |
| `WorldChainMainnet` | World Chain     | 480      | `0x18Bc5bcC660cf2B9cE3cd51a404aFe1a0cBD3C22` |
| `EtherlinkMainnet`  | Etherlink       | 42793    | `0x18bc5bcc660cf2b9ce3cd51a404afe1a0cbd3c22` |
| `GnosisMainnet`     | Gnosis          | 100      | `0x18bc5bcc660cf2b9ce3cd51a404afe1a0cbd3c22` |

### Network Constants

Use the predefined network constants for type safety:

```go
import "github.com/widnyana/idrx-go/blockchain"

// Available network constants
blockchain.BaseMainnet       // "BaseMainnet"
blockchain.PolygonMainnet    // "PolygonMainnet"
blockchain.BSCMainnet        // "BSCMainnet"
blockchain.LiskMainnet       // "LiskMainnet"
blockchain.KaiaMainnet       // "KaiaMainnet"
blockchain.WorldChainMainnet // "WorldChainMainnet"
blockchain.EtherlinkMainnet  // "EtherlinkMainnet"
blockchain.GnosisMainnet     // "GnosisMainnet"
```

### API Design

The blockchain service provides two API patterns:

#### Network Name API (Recommended)

```go
// Type-safe and human-readable
balance, err := client.Blockchain.GetBalanceByNetwork(ctx, blockchain.BaseMainnet, address)
tokenInfo, err := client.Blockchain.GetTokenInfoByNetwork(ctx, blockchain.PolygonMainnet)
```

**Benefits:**

- Type safety with predefined constants
- Human-readable network identifiers
- No need to remember chain IDs
- Consistent naming across environments

#### Chain ID API (Legacy)

```go
// Still supported for backward compatibility
balance, err := client.Blockchain.GetBalance(ctx, 8453, address)
tokenInfo, err := client.Blockchain.GetTokenInfo(ctx, 137)
```

**When to use:**

- Migrating existing code
- Working with dynamic chain IDs
- Interfacing with external systems that use chain IDs

## Architecture

### Contract Layer

- **Generated Bindings**: Auto-generated Go bindings from contract ABIs
- **Proxy Pattern**: Support for ERC-1967 proxy contracts
- **Event Parsing**: Type-safe event log parsing

### Client Layer

- **Multi-Chain Client**: Manages connections to multiple blockchain networks
- **RPC Failover**: Automatic failover between RPC endpoints
- **Transaction Management**: Nonce handling, gas optimization, retry logic

### Service Layer

- **BlockchainService**: High-level API for common operations
- **Integration**: Seamless integration with existing REST API client
- **Error Handling**: Comprehensive error handling and recovery

## Advanced Features

### Gas Estimation

```go
estimate, err := client.Blockchain.EstimateGas(ctx, chainID, "transfer", params)
fmt.Printf("Gas Limit: %d, Gas Price: %s", estimate.GasLimit, estimate.GasPrice)
```

### Platform Fees

```go
fees, err := client.Blockchain.GetPlatformFees(ctx, chainID)
fmt.Printf("Bridge Fee: %d basis points", fees.BurnBridgeFee)
```

### Network Information

```go
networks := client.Blockchain.GetNetworkInfo()
for _, network := range networks {
    fmt.Printf("%s: Chain ID %d", network.Name, network.ChainID)
}
```

### Blacklist Checking

```go
isBlacklisted, err := client.Blockchain.IsAddressBlacklisted(ctx, chainID, address)
```

## Environment Variables

```bash
# Required for blockchain operations
export IDRX_PRIVATE_KEY="0x1234567890abcdef..." # Your private key

# Required for API operations
export IDRX_API_KEY="your-api-key"
export IDRX_SECRET_KEY="your-secret-key"

# Optional: Custom RPC endpoints
export BASE_RPC_URL="https://your-base-rpc.com"
export POLYGON_RPC_URL="https://your-polygon-rpc.com"
```

## Security Considerations

1. **Private Key Management**: Store private keys securely using environment variables or key management services
2. **Gas Limits**: All transactions have configurable gas limits to prevent excessive fees
3. **Input Validation**: All amounts and addresses are validated before blockchain interaction
4. **Blacklist Protection**: Automatic blacklist checking before transactions
5. **Proxy Safety**: Proper handling of proxy contracts to prevent upgrade risks

## Contract Operations

### Token Operations

- `Transfer(to, amount)` - Transfer IDRX tokens
- `BalanceOf(address)` - Get token balance
- `TotalSupply()` - Get total token supply

### Minting Operations (Requires MINTER_ROLE)

- `Mint(to, amount)` - Mint new tokens
- `MintBridge(to, amount, fromChain, nonce)` - Mint from bridge

### Burning Operations

- `Burn(amount)` - Standard token burn
- `BurnWithAccountNumber(amount, accountHash)` - Burn for fiat redemption
- `BurnBridge(amount, toChain)` - Burn for cross-chain bridge

### Bridge Operations

- `GetBridgeNonce()` - Get current bridge nonce
- `IsNonceUsed(fromChain, nonce)` - Check nonce usage
- Bridge event monitoring for cross-chain operations

## Error Handling

```go
// Blockchain operations return detailed errors
result, err := client.Blockchain.Transfer(ctx, chainID, to, amount)
if err != nil {
    switch {
    case strings.Contains(err.Error(), "insufficient funds"):
        // Handle insufficient balance
    case strings.Contains(err.Error(), "blacklisted"):
        // Handle blacklisted address
    case strings.Contains(err.Error(), "gas"):
        // Handle gas-related errors
    default:
        // Handle other errors
    }
}
```

## Development

### Generate Contract Bindings

```bash
make generate-contracts
```

### Run Tests

```bash
make test
```

### Run Example

```bash
go run examples/blockchain_example.go
```

## Integration with REST API

The blockchain service works seamlessly alongside the existing REST API:

```go
// Use REST API for user onboarding and payment processing
mintResp, err := client.Transaction.MintRequest(ctx, &models.MintRequest{
    ToBeMinted: "100000",
    DestinationWalletAddress: walletAddress,
    NetworkChainID: "8453",
})

// Use blockchain service for direct token operations
balance, err := client.Blockchain.GetBalance(ctx, 8453, walletAddress)

// Use REST API for transaction history
history, err := client.Transaction.GetHistory(ctx, &models.TransactionHistoryRequest{
    TransactionType: models.TransactionTypeMint,
})
```

This setup offers the flexibility to use REST APIs for faster, business-level operations while still allowing direct blockchain access when transparency or on-chain verification is needed.
