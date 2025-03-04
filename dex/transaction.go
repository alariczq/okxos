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
)

type GetTransactionStatusRequest struct {
	// Chain Id (e.g., 1 for Ethereum. See Chain IDs), Required
	ChainId string `json:"chainId"`
	// Transaction hash, Required
	TxHash string `json:"txHash"`
	// Set true to check if the transaction is under the current API Key. Set false or omit to query any OKX DEX API transaction.
	IsFromMyProject bool `json:"isFromMyProject"`
}

type TokenDetail struct {
	Symbol       string `json:"symbol"`
	Amount       string `json:"amount"`
	TokenAddress string `json:"tokenAddress"`
}

type TransactionStatusResult struct {
	ChainId          string       `json:"chainId"`
	Hash             string       `json:"hash"`
	Height           string       `json:"height"`
	TxTime           string       `json:"txTime"`
	Status           string       `json:"status"`
	TxType           string       `json:"txType"`
	FromAddress      string       `json:"fromAddress"`
	FromTokenDetails *TokenDetail `json:"fromTokenDetails"`
	ToTokenDetails   *TokenDetail `json:"toTokenDetails"`
	ReferalAmount    string       `json:"referalAmount"`
	ErrorMsg         string       `json:"errorMsg"`
	GasLimit         string       `json:"gasLimit"`
	GasUsed          string       `json:"gasUsed"`
	GasPrice         string       `json:"gasPrice"`
	TxFee            string       `json:"txFee"`
}

// GetTransactionStatus Query the final transaction status of a single-chain swap using txhash.
func (d *DexAPI) GetTransactionStatus(ctx context.Context, req *GetTransactionStatusRequest) (*TransactionStatusResult, error) {
	params := map[string]string{
		"chainId": req.ChainId,
		"txHash":  req.TxHash,
	}
	if req.IsFromMyProject {
		params["isFromMyProject"] = "true"
	}

	var result *TransactionStatusResult
	if err := d.tr.Get(ctx, "/api/v5/dex/aggregator/history", params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
