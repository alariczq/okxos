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

	"github.com/imzhongqi/okxos/errcode"
)

type CrossChainFee struct {
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type CrossChainInfo struct {
	Memo                   string `json:"memo"`
	DestinationChainGasfee string `json:"destinationChainGasfee"`
	DetailStatus           string `json:"detailStatus"`
	Status                 string `json:"status"`
}

type TransactionStatus struct {
	FromChainId        string         `json:"fromChainId"`
	ToChainId          string         `json:"toChainId"`
	FromTxHash         string         `json:"fromTxHash"`
	ToTxHash           string         `json:"toTxHash"`
	FromAmount         string         `json:"fromAmount"`
	FromTokenAddress   string         `json:"fromTokenAddress"`
	ToAmount           string         `json:"toAmount"`
	ToTokenAddress     string         `json:"toTokenAddress"`
	ErrorMsg           string         `json:"errorMsg"`
	BridgeHash         string         `json:"bridgeHash"`
	RefundChainId      string         `json:"refundChainId"`
	RefundTokenAddress string         `json:"refundTokenAddress"`
	RefundTxHash       string         `json:"refundTxHash"`
	SourceChainGasfee  string         `json:"sourceChainGasfee"`
	CrossChainFee      CrossChainFee  `json:"crossChainFee"`
	Symbol             string         `json:"symbol"`
	Address            string         `json:"address"`
	Amount             string         `json:"amount"`
	CrossChainInfo     CrossChainInfo `json:"crossChainInfo"`
}

type GetTransactionStatusRequest struct {
	// Hash address of the source chain
	Hash string `json:"hash"`
	// ChainId Source chain ID (e.g., 1 for Ethereum. See Chain IDs)
	ChainId string `json:"chainId"`
}

// GetTransactionStatus Check the final status of the cross-chain swap according to transaction hash.
func (c *CrossChainAPI) GetTransactionStatus(ctx context.Context, req *GetTransactionStatusRequest) (result *TransactionStatus, err error) {
	params := map[string]string{
		"hash": req.Hash,
	}
	if req.ChainId != "" {
		params["chainId"] = req.ChainId
	}
	var results []*TransactionStatus
	if err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/status", params, &results); err != nil {
		return
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
