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

import "context"

type BridgeInfo struct {
	BridgeName             string   `json:"bridgeName"`
	BridgeId               int64    `json:"bridgeId"`
	RequiredOtherNativeFee bool     `json:"requiredOtherNativeFee"`
	Logo                   string   `json:"logo"`
	SupportedChains        []string `json:"supportedChains"`
}

// GetSupportedBridges Get supported bridges
func (c *CrossChainAPI) GetSupportedBridges(ctx context.Context, chainId string) (result []BridgeInfo, err error) {
	params := map[string]string{}
	if chainId != "" {
		params["chainId"] = chainId
	}
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/supported/bridges", params, &result)
	return
}
