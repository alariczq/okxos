// Copyright (c) 2024-NOW imzhongqi <imzhongqi@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package wallet

import (
	"context"
	"strconv"
	"strings"

	"github.com/imzhongqi/okxos/errcode"
)

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
func (w *WalletAPI) GetTransactionOrder(ctx context.Context, req *TransactionOrderRequest) (result []TransactionOrder, err error) {
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
	var results []TransactionOrder
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
