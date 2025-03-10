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

	"github.com/imzhongqi/okxos/errcode"
)

type ApproveTransactionsRequest struct {
	// ChainId is the chain ID (e.g., 1 for Ethereum. See Chain IDs)
	ChainId string `json:"chainId"`
	// TokenContractAddress is the contract address of a token to be sold (e.g., 0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee)
	TokenContractAddress string `json:"tokenContractAddress"`
	// ApproveAmount is the amount of token that needs to be permitted (set in minimal divisible units,
	// e.g., 1.00 USDT set as 1000000, 1.00 DAI set as 1000000000000000000)
	ApproveAmount string `json:"approveAmount"`
}

type ApproveTransactionsResult struct {
	// Data is the call data
	Data string `json:"data"`
	// DexContractAddress is the contract address of OKX DEX approve (e.g., 0x6f9ffea7370310cd0f890dfde5e0e061059dcfd9)
	DexContractAddress string `json:"dexContractAddress"`
	// GasLimit is the gas limit (e.g., 50000)
	GasLimit string `json:"gasLimit"`
	// GasPrice is the gas price in wei (e.g., 110000000)
	GasPrice string `json:"gasPrice"`
}

// GetApproveTx According to the ERC-20 standard,
// we need to make sure that the OKX router has permission to spend funds with the user's wallet before making a transaction.
//
// This API will generate the relevant data for calling the contract.
func (s *DexAPI) GetApproveTx(ctx context.Context, req *ApproveTransactionsRequest) (*ApproveTransactionsResult, error) {
	params := map[string]string{
		"chainId":              req.ChainId,
		"tokenContractAddress": req.TokenContractAddress,
		"approveAmount":        req.ApproveAmount,
	}
	var results []*ApproveTransactionsResult
	if err := s.tr.Get(ctx, "/api/v5/dex/aggregator/approve-transaction", params, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
