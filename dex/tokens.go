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

package dex

import (
	"context"
	"strconv"
)

type Tokens struct {
	Decimals             string `json:"decimals"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenLogoUrl         string `json:"tokenLogoUrl"`
	TokenName            string `json:"tokenName"`
	TokenSymbol          string `json:"tokenSymbol"`
}

// Get tokens
// It fetches a list of tokens. This interface returns a list of tokens that belong to major platforms or
// are deemed significant enough by OKX. However, you can still quote and swap other tokens outside of this list on OKX DEX.
func (d *DexAPI) GetSupportedTokens(ctx context.Context, chainId int64) (result []Tokens, err error) {
	params := map[string]string{
		"chainId": strconv.FormatInt(chainId, 10),
	}
	err = d.tr.Get(ctx, "/api/v5/dex/aggregator/all-tokens", params, &result)
	return
}
