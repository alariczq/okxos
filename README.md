# OKXOS

This is a Go library for the OKX OS API.

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/imzhongqi/okxos?tab=doc)

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

1. [x] Wallet API
   - [x] Wallet Account Management API
   - [x] Transaction Broadcasting API
2. [x] Dex API
   - [x] Swap API
   - [x] Cross Chain API
   - [ ] Limit Order API **(WIP)**
3. [ ] Marketplace API
4. [ ] DeFi API

## Contributing

Contributions are welcome! Please feel free to submit a PR.

## License

[MIT](LICENSE)
