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

package crosschain

import (
	"context"
	"strconv"
	"strings"

	"github.com/imzhongqi/okxos/errcode"
)

// GetQuoteRequest Get quote request
// Docs: https://www.okx.com/web3/build/docs/waas/dex-get-route-information#get-route-information
type GetQuoteRequest struct {
	FromChainId                     string   `json:"fromChainId"`
	ToChainId                       string   `json:"toChainId"`
	FromTokenAddress                string   `json:"fromTokenAddress"`
	ToTokenAddress                  string   `json:"toTokenAddress"`
	Amount                          string   `json:"amount"`
	Slippage                        string   `json:"slippage"`
	Sort                            int      `json:"sort"`
	FeePercent                      string   `json:"feePercent"`
	AllowBridge                     []string `json:"allowBridge"`
	DenyBridge                      []string `json:"denyBridge"`
	PriceImpactProtectionPercentage string   `json:"priceImpactProtectionPercentage"`
}

type QuoteTokenInfo struct {
	Decimals             int64  `json:"decimals"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenSymbol          string `json:"tokenSymbol"`
}

type DexProtocol struct {
	Percent string `json:"percent"`
	DexName string `json:"dexName"`
}

type SubRouter struct {
	DexProtocol *DexProtocol    `json:"dexProtocol"`
	FromToken   *QuoteTokenInfo `json:"fromToken"`
	ToToken     *QuoteTokenInfo `json:"toToken"`
}

type DexRouter struct {
	Router        string      `json:"router"`
	RouterPercent string      `json:"routerPercent"`
	SubRouterList []SubRouter `json:"subRouterList"`
}

type BridgeRouter struct {
	BridgeId                  int    `json:"bridgeId"`
	BridgeName                string `json:"bridgeName"`
	CrossChainFee             string `json:"crossChainFee"`
	CrossChainFeeTokenAddress string `json:"crossChainFeeTokenAddress"`
	OtherNativeFee            string `json:"otherNativeFee"`
	EstimateGasFee            string `json:"estimateGasFee"`
	EstimatedTime             string `json:"estimatedTime"`
}

type Router struct {
	EstimateTime        string        `json:"estimateTime"`
	EstimateGasFee      string        `json:"estimateGasFee"`
	FromChainNetworkFee string        `json:"fromChainNetworkFee"`
	ToChainNetworkFee   string        `json:"toChainNetworkFee"`
	ToTokenAmount       string        `json:"toTokenAmount"`
	MinimumReceived     string        `json:"minimumReceived"`
	NeedApprove         int           `json:"needApprove"`
	Router              *BridgeRouter `json:"router"`
	FromDexRouterList   []DexRouter   `json:"fromDexRouterList"`
	ToDexRouterList     []DexRouter   `json:"toDexRouterList"`
}

type QuoteResult struct {
	FromChainId     string         `json:"fromChainId"`
	ToChainId       string         `json:"toChainId"`
	FromTokenAmount string         `json:"fromTokenAmount"`
	FromToken       QuoteTokenInfo `json:"fromToken"`
	ToToken         QuoteTokenInfo `json:"toToken"`
	RouterList      []Router       `json:"routerList"`
}

// GetQuote Get quote
func (c *CrossChainAPI) GetQuote(ctx context.Context, quote *GetQuoteRequest) (result *QuoteResult, err error) {
	params := map[string]string{
		"fromChainId":      quote.FromChainId,
		"toChainId":        quote.ToChainId,
		"fromTokenAddress": quote.FromTokenAddress,
		"toTokenAddress":   quote.ToTokenAddress,
		"amount":           quote.Amount,
		"slippage":         quote.Slippage,
	}
	if quote.Sort != 0 {
		params["sort"] = strconv.Itoa(quote.Sort)
	}
	if quote.FeePercent != "" {
		params["feePercent"] = quote.FeePercent
	}
	if len(quote.AllowBridge) > 0 {
		params["allowBridge"] = strings.Join(quote.AllowBridge, ",")
	}
	if len(quote.DenyBridge) > 0 {
		params["denyBridge"] = strings.Join(quote.DenyBridge, ",")
	}
	if quote.PriceImpactProtectionPercentage != "" {
		params["priceImpactProtectionPercentage"] = quote.PriceImpactProtectionPercentage
	}

	var results []*QuoteResult
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/quote", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
