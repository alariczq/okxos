package wallet

import (
	"context"
	"strconv"
	"strings"

	"github.com/imzhongqi/okxos/errcode"
)

type ValidateAddressRequest struct {
	// Unique identifier for the chain
	ChainIndex string `json:"chainIndex"`
	// Address
	Address string `json:"address"`
}

// AddressType
// 0: Invalid address format
// 1: Valid user address
// 2: Valid contract address
type AddressType string

func (a AddressType) IsInvalidAddress() bool {
	return a == "0"
}

func (a AddressType) IsUserAddress() bool {
	return a == "1"
}

func (a AddressType) IsContractAddress() bool {
	return a == "2"
}

type ValidateAddressResult struct {
	AddressType AddressType `json:"addressType"`
	// Whether the address has hit blacklist check
	HitBlacklist bool `json:"hitBlacklist"`
	// Tag type for blacklisted addresses, including phishing, contract vulnerabilities, etc.
	Tag string `json:"tag"`
}

// ValidateAddress Provide an address to determine if it is a valid user or contract address, and whether it has hit blacklist check.
func (w *WalletAPI) ValidateAddress(ctx context.Context, req *ValidateAddressRequest) (result *ValidateAddressResult, err error) {
	params := map[string]string{
		"chainIndex": req.ChainIndex,
		"address":    req.Address,
	}
	var results []*ValidateAddressResult
	err = w.tr.Get(ctx, "/api/v5/wallet/pre-transaction/validate-address", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type TransactionBroadcastRequest struct {
	SignedTx   string `json:"signedTx"`
	ChainIndex string `json:"chainIndex"`
	Address    string `json:"address"`
	AccountId  string `json:"accountId"`
}

type TransactionBroadcastResult struct {
	OrderId string `json:"orderId"`
}

func (w *WalletAPI) TransactionBroadcast(ctx context.Context, tx *TransactionBroadcastRequest) (result *TransactionBroadcastResult, err error) {
	params := map[string]string{
		"signedTx":   tx.SignedTx,
		"chainIndex": tx.ChainIndex,
		"address":    tx.Address,
	}
	if tx.AccountId != "" {
		params["accountId"] = tx.AccountId
	}
	err = w.tr.Post(ctx, "/api/v5/wallet/pre-transaction/broadcast-transaction", params, &result)
	return
}

type TransactionOrderRequest struct {
	Address    string `json:"address"`
	AccountId  string `json:"accountId"`
	ChainIndex string `json:"chainIndex"`
	// Transaction Status
	// 1: Pending
	// 2: Success
	// 3: Failed
	TxStatus string `json:"txStatus"`
	OrderId  string `json:"orderId"`
	Cursor   string `json:"cursor"`
	Limit    string `json:"limit"`
}

type TransactionOrder struct {
	ChainIndex int64  `json:"chainIndex"`
	Address    string `json:"address"`
	AccountId  string `json:"accountId"`
	OrderId    string `json:"orderId"`
	TxStatus   string `json:"txStatus"`
	TxHash     string `json:"txHash"`
	Limit      string `json:"limit"`
}

// Get Transaction Order
func (w *WalletAPI) GetTransactionOrder(ctx context.Context, req *TransactionOrderRequest) (result []*TransactionOrder, err error) {
	params := map[string]string{}
	if req.Address != "" {
		params["address"] = req.Address
	}
	if req.AccountId != "" {
		params["accountId"] = req.AccountId
	}
	if req.ChainIndex != "" {
		params["chainIndex"] = req.ChainIndex
	}
	if req.TxStatus != "" {
		params["txStatus"] = req.TxStatus
	}
	if req.OrderId != "" {
		params["orderId"] = req.OrderId
	}
	if req.Cursor != "" {
		params["cursor"] = req.Cursor
	}
	if req.Limit != "" {
		params["limit"] = req.Limit
	}
	var results []*TransactionOrder
	err = w.tr.Get(ctx, "/api/v5/wallet/post-transaction/orders", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results, nil
}

type GetTransactionHistoryByAddressRequest struct {
	// Address to query the transaction history for
	Address string   `json:"address"`
	Chains  []string `json:"chains"`
	// Unique identifier for the chain, e.g. ETH=3
	ChainIndex string `json:"chainIndex"`
	// Token contract address; if empty, query addresses with main chain currency balance;if not pass, query all
	TokenAddress string `json:"tokenAddress"`
	// Start time, queries transactions after this time. Unix timestamp, in milliseconds, e.g., 1597026383085
	Begin string `json:"begin"`
	// End time, queries transactions before this time. If both begin and end are not provided, queries transactions before the current time. Unix timestamp, in milliseconds, e.g., 1597026383085
	End string `json:"end"`
	// Cursor
	Cursor string `json:"cursor"`
	// Number of records to return, defaults to the most recent 20 records.
	Limit string `json:"limit"`
	// Option to filter out potentially risky airdrop tokens.
	// Defaults to filtering.
	// true: filter
	// false: do not filter
	ExcludeRiskToken *bool `json:"excludeRiskToken"`
}

type AddressBalance struct {
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type TransactionHistory struct {
	// Unique identifier for the chain, e.g. ETH=3
	ChainIndex string `json:"chainIndex"`
	// Transaction hash
	TxHash string `json:"txHash"`
	// EVM Transaction tier type
	// 0: Outer main chain coin transfer
	// 1: Contract inner main chain coin transfer
	// 2: Token transfer
	IType string `json:"iType"`
	// Contract Function Call
	MethodId string `json:"methodId"`
	// The nth transaction initiated by the sender address
	Nonce string `json:"nonce"`
	// Transaction time in Unix timestamp format, in milliseconds, e.g., 1597026383085
	TxTime string `json:"txTime"`
	// Sending/input address, comma-separated for multi-signature transactions
	From []*AddressBalance `json:"from"`
	// Receiving/output address, comma-separated for multiple addresses
	To []*AddressBalance `json:"to"`
	// Token contract address
	TokenAddress string `json:"tokenAddress"`
	// Transaction amount
	Amount string `json:"amount"`
	// Currency symbol corresponding to the transaction amount
	Symbol string `json:"symbol"`
	// Transaction fee
	TxFee string `json:"txFee"`
	// Transaction status: success for successful transactions, fail for failed transactions, pending for pending transactions
	TxStatus string `json:"txStatus"`
	// false: Not in blacklist, true: In blacklist
	HitBlacklist bool `json:"hitBlacklist"`
	// Tag type for blacklisted addresses, including phishing, contract vulnerabilities, etc.
	Tag string `json:"tag"`
}

type GetTransactionHistoryByAddressResult struct {
	// List of transactions
	TransactionList []*TransactionHistory `json:"transactionList"`
	// Cursor
	Cursor string `json:"cursor"`
}

// Get Transaction History By Address
func (w *WalletAPI) GetTransactionHistoryByAddress(ctx context.Context, req *GetTransactionHistoryByAddressRequest) (result *GetTransactionHistoryByAddressResult, err error) {
	params := map[string]string{
		"address": req.Address,
	}
	if len(req.Chains) > 0 {
		params["chains"] = strings.Join(req.Chains, ",")
	}
	if req.TokenAddress != "" {
		params["tokenAddress"] = req.TokenAddress
	}
	if req.Begin != "" {
		params["begin"] = req.Begin
	}
	if req.End != "" {
		params["end"] = req.End
	}
	if req.Cursor != "" {
		params["cursor"] = req.Cursor
	}
	if req.Limit != "" {
		params["limit"] = req.Limit
	}
	if req.ExcludeRiskToken != nil {
		params["excludeRiskToken"] = strconv.FormatBool(*req.ExcludeRiskToken)
	}
	var results []*GetTransactionHistoryByAddressResult
	err = w.tr.Get(ctx, "/api/v5/wallet/post-transaction/transactions-by-address", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
