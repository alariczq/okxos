# Go OKX OS API

This is a Go library for the OKX OS API.

## Installation

```bash
go get -u github.com/imzhongqi/okxos
```

## Example

```go
client := client.NewClient("key", "secret", "passphrase",
    client.WithProjectID("test"),
)

walletapi := wallet.NewWalletAPI(client)

_, err := walletapi.ListSupportedChains(context.Background())
if err != nil {
    panic(err)
}
```

# Roadmap

1. [ ] Wallet API
   - [x] Wallet Account Management API
   - [x] Transaction Broadcasting API
2. [ ] Dex API
   - [x] Swap API
   - [ ] Cross Chain API
3. [ ] Marketplace API
4. [ ] DeFi API

# License

[MIT](LICENSE)
