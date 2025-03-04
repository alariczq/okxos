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

	"github.com/imzhongqi/okxos/errcode"
)

type SupportedChains struct {
	Name       string `json:"name"`
	LogoUrl    string `json:"logoUrl"`
	ShortName  string `json:"shortName"`
	ChainIndex string `json:"chainIndex"`
}

// Supported Blockchains
func (w *WalletAPI) SupportedChains(ctx context.Context) (result []*SupportedChains, err error) {
	err = w.tr.Get(ctx, "/api/v5/wallet/chain/supported-chains", nil, &result)
	return
}

type TokenIndexPriceRequest struct {
	ChainIndex   string `json:"chainIndex"`
	TokenAddress string `json:"tokenAddress"`
}

type TokenPrice struct {
	ChainIndex   string `json:"chainIndex"`
	TokenAddress string `json:"tokenAddress"`
	Price        string `json:"price"`
	Time         string `json:"time"`
}

func (w *WalletAPI) TokenIndexPrice(ctx context.Context, req []*TokenIndexPriceRequest) (result []TokenPrice, err error) {
	err = w.tr.Post(ctx, "/api/v5/wallet/token/current-price", req, &result)
	return
}

type RealTimeTokenPriceRequest struct {
	ChainIndex   string `json:"chainIndex"`
	TokenAddress string `json:"tokenAddress"`
}

// Real-time Token Price
func (w *WalletAPI) RealTimeTokenPrice(ctx context.Context, req []*RealTimeTokenPriceRequest) (result []TokenPrice, err error) {
	err = w.tr.Post(ctx, "/api/v5/wallet/token/real-time-price", req, &result)
	return
}

type HistoricalTokenPriceRequest struct {
	ChainIndex string `json:"chainIndex"`
	// Token address.
	// 1: Pass an empty string "" to query the native token of the corresponding chain.
	// 2: Pass the specific token contract address to query the corresponding token.
	// 3: Different inscription tokens are passed in the following formats:
	// FBRC-20: Use fbtc_fbrc20_name, such as fbtc_fbrc20_babymusk
	// BRC-20: Use btc-brc20-tick(name), such as btc-brc20-ordi
	// Runes: Use btc-runesMain-tickId, such as btc-runesMain-840000:2
	// SRC-20: Use btc-src20-name, such as btc-src20-utxo
	TokenAddress string `json:"tokenAddress"`
	// Number of entries per query, default is 50, maximum is 200
	Limit int64 `json:"limit"`
	// Cursor position, defaults to the first entry
	Cursor int64 `json:"cursor"`
	// Start time to query historical prices after. Unix timestamp in milliseconds
	Begin int64 `json:"begin"`
	// End time to query historical prices before.
	// If neither begin nor end is provided, query historical prices before the current time.
	// Unix timestamp in milliseconds
	End int64 `json:"end"`
	// Time interval unit:
	// 1m: 1 minute
	// 5m: 5 minutes
	// 30m: 30 minutes
	// 1h: 1 hour
	// 1d: 1 day (default)
	Period string `json:"period"`
}

type HistoricalTokenPriceResult struct {
	Cursor string        `json:"cursor"`
	Prices []*TokenPrice `json:"prices"`
}

// Historical Token Price
func (w *WalletAPI) HistoricalTokenPrice(ctx context.Context, req *HistoricalTokenPriceRequest) (result *HistoricalTokenPriceResult, err error) {
	params := map[string]string{
		"chainIndex":   req.ChainIndex,
		"tokenAddress": req.TokenAddress,
		"limit":        strconv.FormatInt(req.Limit, 10),
		"cursor":       strconv.FormatInt(req.Cursor, 10),
		"begin":        strconv.FormatInt(req.Begin, 10),
		"end":          strconv.FormatInt(req.End, 10),
		"period":       req.Period,
	}
	err = w.tr.Get(ctx, "/api/v5/wallet/token/historical-price", params, &result)
	return
}

type ProjectInformationRequest struct {
	ChainIndex   string `json:"chainIndex"`
	TokenAddress string `json:"tokenAddress"`
}

type SocialUrls struct {
	Messageboard []string `json:"messageboard"`
	Github       []string `json:"github"`
	Twitter      []string `json:"twitter"`
	Chat         []string `json:"chat"`
	Reddit       []string `json:"reddit"`
}

type ProjectInformation struct {
	LogoUrl         string     `json:"logoUrl"`
	OfficialWebsite string     `json:"officialWebsite"`
	SocialUrls      SocialUrls `json:"socialUrls"`
	Decimals        string     `json:"decimals"`
	TokenAddress    string     `json:"tokenAddress"`
	ChainIndex      string     `json:"chainIndex"`
	ChainName       string     `json:"chainName"`
	Symbol          string     `json:"symbol"`
	MaxSupply       string     `json:"maxSupply"`
	TotalSupply     string     `json:"totalSupply"`
	Volume24h       string     `json:"volume24h"`
	MarketCap       string     `json:"marketCap"`
}

// Project Information
func (w *WalletAPI) ProjectInformation(ctx context.Context, req *ProjectInformationRequest) (result *ProjectInformation, err error) {
	params := map[string]string{
		"chainIndex":   req.ChainIndex,
		"tokenAddress": req.TokenAddress,
	}
	var results []*ProjectInformation
	err = w.tr.Get(ctx, "/api/v5/wallet/token/token-detail", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
