# IDRX Go SDK Examples

This directory contains practical examples demonstrating core IDRX SDK functionality.

## Prerequisites

Set the following environment variables:

```bash
export IDRX_API_KEY="your-api-key"
export IDRX_API_SECRET="your-api-secret"
```

For blockchain operations, you'll also need:

```bash
export WALLET_PRIVATE_KEY="your-wallet-private-key-hex"
```

## Available Examples

### 1. User Onboarding (`onboarding/`)

Demonstrates business authentication and user KYC onboarding workflow.

```bash
cd onboarding
go run main.go
```

**Shows:**
- Client initialization with business credentials
- User onboarding with KYC data
- File upload handling
- Extracting user credentials

### 2. Mint & Redemption (`mint-redemption/`)

Complete workflow for minting IDRX and redeeming to bank account.

```bash
cd mint-redemption
go run main.go
```

**Shows:**
- User authentication
- Querying swap rates
- Creating mint requests
- Adding bank accounts
- Submitting redemption requests

### 3. Blacklist Check (`blacklist-check/`)

Blockchain operations for checking wallet blacklist status.

```bash
cd blacklist-check
go run main.go
```

**Shows:**
- Blockchain service initialization
- Multi-network blacklist verification
- Network configuration

## Running Examples

Each example is a standalone program. Navigate to the example directory and run:

```bash
go run main.go
```

## Documentation

For complete API documentation, visit:
- [IDRX Documentation](https://docs.idrx.co)
- [SDK README](../README.md)

## Authentication Context

- **Business Auth** (IDRX_API_KEY + IDRX_API_SECRET): Organization management, user onboarding
- **User Auth**: Individual transactions - credentials obtained from onboarding response
- **Blockchain Auth** (WALLET_PRIVATE_KEY): Direct on-chain operations

## Support

For issues or questions, please refer to the main repository documentation.
