# IDRX Go SDK

Go SDK for [IDRX](https://docs.idrx.co) — a regulated stablecoin system for fiat on/off-ramp, multi-chain operations, and dual-auth KYC flows.

```sh
go get github.com/widnyana/idrx-go
```

## What You Get

- Dual authentication (Business & User).
- Full REST API and Smart Contract coverage
- Multi-chain support: Base, Polygon, BSC, Lisk, Kaia, World Chain, Etherlink, Gnosis
- Safe production use: strong type validation, context, retries, structured error.

---

## Authentication

```go
// Business context
biz := idrx.NewClient(idrx.WithBusinessAuth(bizKey, bizSecret))

// User context (from onboarding response)
user := idrx.NewClient(idrx.WithUserAuth(userKey, userSecret))
```

---

## Common Operations

### User Onboarding (Business Auth)

```go
resp, err := biz.Account.Onboard(ctx, &models.OnboardingRequest{
    Email:    "user@company.com",
    Fullname: "John Doe",
})
userKey, userSecret := resp.APIKey, resp.APISecret
```

### Mint IDRX (User Auth)

```go
resp, err := user.Transaction.MintRequest(ctx, &models.MintRequest{
    ToBeMinted: "100000",
    NetworkChainID: blockchain.PolygonChainID, // Polygon Mainnet
})
// redirect user to resp.PaymentURL
```

### Redeem IDRX (User Auth)

```go
resp, err := user.Transaction.RedeemRequest(ctx, &models.RedeemRequest{
    TxHash: "0x1234...", BankCode: "BNK",
})
```

### Bank Accounts

```go
acc, _ := user.Account.AddBankAccount(ctx, &models.AddBankAccountRequest{BankCode: "BNK"})
list, _ := user.Account.GetBankAccounts(ctx)
_ = user.Account.DeleteBankAccount(ctx, acc.ID)
```

---

## Advanced

Custom client:

```go
client := idrx.NewClient(
    idrx.WithBusinessAuth(apiKey, secret),
    idrx.WithTimeout(45*time.Second),
)
```

Error handling:

```go
if err != nil {
    if e, ok := err.(*models.APIError); ok {
        log.Printf("[%d] %s", e.StatusCode, e.Message)
    }
}
```

For detailed blockchain integration, see [BLOCKCHAIN.md](BLOCKCHAIN.md).

---

## License

[MIT](LICENSE)

For detailed APIs and blockchain setup, see [`examples/README.md`](examples/README.md) or [docs.idrx.co](https://docs.idrx.co).
