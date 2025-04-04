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

type GetTotalValueByAddressRequest struct {
	// Address Get the total valuation for the address
	Address string `json:"address"`
	// Chains Query the total balance of the multiple chains, which can be separated by ",". Supports up to 50 chains
	Chains []string `json:"chains"`
	// AssetType Type of asset to query, defaults to total balance of all assets.
	// 0: Query total balance of all assets, including tokens and DeFi assets;
	// 1: Query only token balance;
	// 2: Query only DeFi balance
	AssetType string `json:"assetType"`
	// ExcludeRiskToken Option to filter out potentially risky airdrop tokens.
	// Defaults to filtering.
	// true: filter
	// false: do not filter
	ExcludeRiskToken *bool `json:"excludeRiskToken"`
}

type TotalValueResult struct {
	// Returns the total asset balance based on the query asset type, expressed in USD
	TotalValue string `json:"totalValue"`
}

// Get Total Value By Address
func (w *WalletAPI) GetTotalValueByAddress(ctx context.Context, req *GetTotalValueByAddressRequest) (result *TotalValueResult, err error) {
	params := map[string]string{
		"address":   req.Address,
		"chains":    strings.Join(req.Chains, ","),
		"assetType": req.AssetType,
	}
	if req.ExcludeRiskToken != nil {
		params["excludeRiskToken"] = strconv.FormatBool(*req.ExcludeRiskToken)
	}
	var results []*TotalValueResult
	err = w.tr.Get(ctx, "/api/v5/wallet/asset/total-value-by-address", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetTotalTokenBalancesByAddressRequest struct {
	// Address Query the token balances for the specified address
	Address string `json:"address"`
	// Chains Query the token balances for the specified chains, which can be separated by ",". A maximum of 50 chains is supported.
	Chains []string `json:"chains"`
	// Filter
	// 0: Filter out risk airdrop tokens
	// 1: Do not filter
	// Default is to filter
	Filter string `json:"filter"`
}

type TokenBalance struct {
	ChainIndex      string `json:"chainIndex"`
	TokenAddress    string `json:"tokenAddress"`
	Address         string `json:"address"`
	Symbol          string `json:"symbol"`
	Balance         string `json:"balance"`
	TokenPrice      string `json:"tokenPrice"`
	TokenType       string `json:"tokenType"`
	TransferAmount  string `json:"transferAmount"`
	AvailableAmount string `json:"availableAmount"`
	IsRiskToken     bool   `json:"isRiskToken"`
}

type TokenBalanceResult struct {
	TokenAssets []*TokenBalance `json:"tokenAssets"`
	TimeStamp   string          `json:"timeStamp"`
}

// Get Total Token Balances By Address
func (w *WalletAPI) GetTotalTokenBalancesByAddress(ctx context.Context, req *GetTotalTokenBalancesByAddressRequest) (result *TokenBalanceResult, err error) {
	params := map[string]string{
		"address": req.Address,
		"chains":  strings.Join(req.Chains, ","),
		"filter":  req.Filter,
	}
	var results []*TokenBalanceResult
	err = w.tr.Get(ctx, "/api/v5/wallet/asset/all-token-balances-by-address", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type TokenAddress struct {
	TokenAddress string `json:"tokenAddress"`
	ChainIndex   string `json:"chainIndex"`
}

type GetTokenBalancesByAddressRequest struct {
	// Address Query the token balances for the specified address
	Address string `json:"address"`
	// TokenAddresses Query the token balances for the specified token address
	TokenAddresses []*TokenAddress `json:"tokenAddresses"`
	// Filter
	// 0: Filter out risk airdrop tokens
	// 1: Do not filter
	// Default is to filter
	Filter string `json:"filter"`
}

// Get Token Balances By Address
func (w *WalletAPI) GetTokenBalancesByAddress(ctx context.Context, req *GetTokenBalancesByAddressRequest) (result *TokenBalanceResult, err error) {
	var results []*TokenBalanceResult
	err = w.tr.Post(ctx, "/api/v5/wallet/asset/token-balances-by-address", req, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetTotalValueByAccountRequest struct {
	// AccountId Query the total valuation for the specified account
	AccountId string `json:"accountId"`
	// Chains Query the total valuation for the specified chains, which can be separated by ",". A maximum of 50 chains is supported.
	Chains []string `json:"chains"`
	// AssetType Type of asset to query, defaults to total balance of all assets.
	// 0: Query total balance of all assets, including tokens and DeFi assets;
	// 1: Query only token balance;
	// 2: Query only DeFi balance
	AssetType string `json:"assetType"`
	// ExcludeRiskToken Option to filter out potentially risky airdrop tokens. Defaults to filtering.
	// true: filter, false: do not filter
	ExcludeRiskToken *bool `json:"excludeRiskToken"`
}

// Get Total Value By Account
func (w *WalletAPI) GetTotalValueByAccount(ctx context.Context, req *GetTotalValueByAccountRequest) (result *TotalValueResult, err error) {
	params := map[string]string{
		"accountId": req.AccountId,
		"chains":    strings.Join(req.Chains, ","),
		"assetType": req.AssetType,
	}
	if req.ExcludeRiskToken != nil {
		params["excludeRiskToken"] = strconv.FormatBool(*req.ExcludeRiskToken)
	}
	var results []*TotalValueResult
	err = w.tr.Get(ctx, "/api/v5/wallet/asset/total-value", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetTotalTokenBalancesByAccountRequest struct {
	// AccountId Query the token balances for the specified account
	AccountId string `json:"accountId"`
	// Chains Query the token balances for the specified chains, which can be separated by ",". A maximum of 50 chains is supported.
	Chains []string `json:"chains"`
	// Filter
	// 0: Filter out risk airdrop tokens
	// 1: Do not filter
	// Default is to filter
	Filter string `json:"filter"`
}

// Get Total Token Balances By Account
func (w *WalletAPI) GetTotalTokenBalancesByAccount(ctx context.Context, req *GetTotalTokenBalancesByAccountRequest) (result *TokenBalanceResult, err error) {
	params := map[string]string{
		"accountId": req.AccountId,
	}
	if len(req.Chains) > 0 {
		params["chains"] = strings.Join(req.Chains, ",")
	}
	if req.Filter != "" {
		params["filter"] = req.Filter
	}
	var results []*TokenBalanceResult
	err = w.tr.Get(ctx, "/api/v5/wallet/asset/wallet-all-token-balances", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetTokenBalancesByAccountRequest struct {
	// AccountId Query the token balances for the specified account
	AccountId string `json:"accountId"`
	// TokenAddresses Query the token balances for the specified token address
	TokenAddresses []*TokenAddress `json:"tokenAddresses"`
}

// Get Token Balances By Account
func (w *WalletAPI) GetTokenBalancesByAccount(ctx context.Context, req *GetTokenBalancesByAccountRequest) (result *TokenBalanceResult, err error) {
	var results []*TokenBalanceResult
	err = w.tr.Post(ctx, "/api/v5/wallet/asset/token-balances", req, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
