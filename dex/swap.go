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
	"strings"

	"github.com/imzhongqi/okxos/errcode"
)

type Tx struct {
	Data                 string   `json:"data"`
	From                 string   `json:"from"`
	Gas                  string   `json:"gas"`
	GasPrice             string   `json:"gasPrice"`
	MaxPriorityFeePerGas string   `json:"maxPriorityFeePerGas"`
	MinReceiveAmount     string   `json:"minReceiveAmount"`
	SignatureData        []string `json:"signatureData"`
	To                   string   `json:"to"`
	Value                string   `json:"value"`
}

// GetSwapTxRequest
type GetSwapTxRequest struct {
	// Chain Id (e.g., 1 for Ethereum. See Chain IDs), Required
	ChainId string `json:"chainId"`
	// The input amount of a token to be sold, Required
	Amount string `json:"amount"`
	// The contract address of a token you want to send (e.g.,0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee), Required
	FromTokenAddress string `json:"fromTokenAddress"`
	// The contract address of a token you want to receive (e.g.,0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48), Required
	ToTokenAddress string `json:"toTokenAddress"`
	// The slippage you are willing to accept. If you set 0.5, it means 50% slippage is acceptable. min:0 max:1, Required
	Slippage string `json:"slippage"`
	// User's wallet address (e.g.,0x3f6a3f57569358a512ccc0e513f171516b0fd42a), Required
	UserWalletAddress string `json:"userWalletAddress"`
	// ReferrerAddress (Supports SOL or SPL Token commissions.
	// supports SOL using wallet address or SPL token commissions using token account)
	// The fromToken address that receives the commission.
	// When using the API, you need to configure the commission ratio using feePercent.
	// Each transaction can only choose commission from either the fromToken or the toToken.
	// Note:
	// 1. For EVM chains: Transactions involving wrapped pairs, such as ETH and WETH, are not supported here.
	// 2. For Solana chain: The commission address must have some SOL deposited in advance for activation.
	ReferrerAddress string `json:"referrerAddress"`
	// Recipient address of a purchased token if not set,
	//userWalletAddress will receive a purchased token (e.g.,0x3f6a3f57569358a512ccc0e513f171516b0fd42a)
	SwapReceiverAddress string `json:"swapReceiverAddress"`
	// The percentage of fromTokenAmount will be sent to the referrer's address,
	// the rest will be set as the input amount to be sold.
	// Min percentage: 0. Max percentage: 3. Maximum 2 decimal points.
	// Longer sections will be automatically omitted. (E.g. 1.326% is the actual input, but the final calculation will only adopt 1.32%.)
	FeePercent string `json:"feePercent"`
	// The gas (in wei) for the swap transaction. If the value is too low to achieve the quote, an error will be returned
	Gaslimit string `json:"gaslimit"`
	// The target gas price level for the swap transaction,set to average or fast or slow
	GasLevel string `json:"gasLevel"`
	// DexId of the liquidity pool for limited quotes, multiple combinations separated by , (e.g., 1,50,180, see liquidity list for more)
	DexIds []string `json:"dexIds"`
	// The percentage (between 0 - 1.0) of the price impact allowed.
	PriceImpactProtectionPercentage string `json:"priceImpactProtectionPercentage"`
	// You can customize the parameters to be sent on the blockchain in callData by encoding the
	// data into a 128-character 64-bytes hexadecimal string.
	// For example, the string
	// "0x111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	// needs to keep the "0x" at its start.
	CallDataMemo string `json:"callDataMemo"`
	// The toToken address that receives the commission.
	ToTokenReferrerAddress string `json:"toTokenReferrerAddress"`
	// Used for transactions on the Solana network and similar to gasPrice on Ethereum.
	// This price determines the priority level of the transaction.
	// The higher the price, the more likely that the transaction can be processed faster.
	ComputeUnitPrice string `json:"computeUnitPrice"`
	// Used for transactions on the Solana network and analogous to gasLimit on Ethereum,
	// which ensures that the transaction won't take too much computing resource.
	ComputeUnitLimit string `json:"computeUnitLimit"`
	// The wallet address to receive the commission fee from the fromToken.
	// This new field no longer requires a token account parameter for SPL Token;
	// specifying the Sol wallet address is sufficient.
	FromTokenReferrerWalletAddress string `json:"fromTokenReferrerWalletAddress"`
	// The wallet address to receive the commission fee from the toToken.
	// This new field no longer requires a token account parameter for SPL-Token.
	ToTokenReferrerWalletAddress string `json:"toTokenReferrerWalletAddress"`
	// Default is false. When set to true, the original slippage (if set) will be covered by the autoSlippage and
	// the API will calculate and return auto slippage recommendations based on current market data.
	AutoSlippage bool `json:"autoSlippage"`
	// When autoSlippage is set to true, this value is the maximum auto slippage returned by the API.
	// We recommend that users adopt this value to ensure risk control.
	MaxAutoSlippage string `json:"maxAutoSlippage"`
}

type GetSwapTxResult struct {
	RouterResult *QuotesResult `json:"routerResult"`
	Tx           *Tx           `json:"tx"`
}

// Swap generates the data to call the OKX DEX router to execute a swap.
// https://www.okx.com/web3/build/docs/waas/dex-swap
// Note:
// 1. For EVM chains: Transactions involving wrapped pairs, such as ETH and WETH, are not supported here.
// 2. For Solana chain: The commission address must have some SOL deposited in advance for activation.
func (d *DexAPI) GetSwapTx(ctx context.Context, swap *GetSwapTxRequest) (result *GetSwapTxResult, err error) {
	params := map[string]string{
		"chainId":           swap.ChainId,
		"amount":            swap.Amount,
		"fromTokenAddress":  swap.FromTokenAddress,
		"toTokenAddress":    swap.ToTokenAddress,
		"slippage":          swap.Slippage,
		"userWalletAddress": swap.UserWalletAddress,
	}
	if swap.ReferrerAddress != "" {
		params["referrerAddress"] = swap.ReferrerAddress
	}
	if swap.SwapReceiverAddress != "" {
		params["swapReceiverAddress"] = swap.SwapReceiverAddress
	}
	if swap.FeePercent != "" {
		params["feePercent"] = swap.FeePercent
	}
	if swap.Gaslimit != "" {
		params["gaslimit"] = swap.Gaslimit
	}
	if swap.GasLevel != "" {
		params["gasLevel"] = swap.GasLevel
	}
	if len(swap.DexIds) > 0 {
		params["dexIds"] = strings.Join(swap.DexIds, ",")
	}
	if swap.PriceImpactProtectionPercentage != "" {
		params["priceImpactProtectionPercentage"] = swap.PriceImpactProtectionPercentage
	}
	if swap.CallDataMemo != "" {
		params["callDataMemo"] = swap.CallDataMemo
	}
	if swap.ToTokenReferrerAddress != "" {
		params["toTokenReferrerAddress"] = swap.ToTokenReferrerAddress
	}
	if swap.ComputeUnitPrice != "" {
		params["computeUnitPrice"] = swap.ComputeUnitPrice
	}
	if swap.ComputeUnitLimit != "" {
		params["computeUnitLimit"] = swap.ComputeUnitLimit
	}
	if swap.FromTokenReferrerWalletAddress != "" {
		params["fromTokenReferrerWalletAddress"] = swap.FromTokenReferrerWalletAddress
	}
	if swap.ToTokenReferrerWalletAddress != "" {
		params["toTokenReferrerWalletAddress"] = swap.ToTokenReferrerWalletAddress
	}
	if swap.AutoSlippage {
		params["autoSlippage"] = "true"
	}
	if swap.MaxAutoSlippage != "" {
		params["maxAutoSlippage"] = swap.MaxAutoSlippage
	}

	var results []*GetSwapTxResult
	if err = d.tr.Get(ctx, "/api/v5/dex/aggregator/swap", params, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}

	return results[0], nil
}

type GetSolSwapInstructionRequest struct {
	ChainId                         string
	Amount                          string
	FromTokenAddress                string
	ToTokenAddress                  string
	Slippage                        string
	UserWalletAddress               string
	SwapReceiverAddress             string
	FeePercent                      string
	FromTokenReferrerWalletAddress  string
	ToTokenReferrerWalletAddress    string
	DexIds                          []string
	PriceImpactProtectionPercentage string
	ComputeUnitPrice                string
	ComputeUnitLimit                string
}

type GetSolSwapInstructionResult struct {
	AddressLookupTableAccount []string          `json:"addressLookupTableAccount"`
	InstructionLists          []InstructionInfo `json:"instructionLists"`
}

type InstructionInfo struct {
	Data      string        `json:"data"`
	Accounts  []AccountInfo `json:"accounts"`
	ProgramId string        `json:"programId"`
}

type AccountInfo struct {
	IsSigner   bool   `json:"isSigner"`
	IsWritable bool   `json:"isWritable"`
	Pubkey     string `json:"pubkey"`
}

func (d *DexAPI) GetSolSwapInstruction(ctx context.Context, req *GetSolSwapInstructionRequest) (result *GetSolSwapInstructionResult, err error) {
	params := map[string]string{
		"chainId":           req.ChainId,
		"amount":            req.Amount,
		"fromTokenAddress":  req.FromTokenAddress,
		"toTokenAddress":    req.ToTokenAddress,
		"slippage":          req.Slippage,
		"userWalletAddress": req.UserWalletAddress,
	}
	if req.SwapReceiverAddress != "" {
		params["swapReceiverAddress"] = req.SwapReceiverAddress
	}
	if req.FeePercent != "" {
		params["feePercent"] = req.FeePercent
	}
	if req.FromTokenReferrerWalletAddress != "" {
		params["fromTokenReferrerWalletAddress"] = req.FromTokenReferrerWalletAddress
	}
	if req.ToTokenReferrerWalletAddress != "" {
		params["toTokenReferrerWalletAddress"] = req.ToTokenReferrerWalletAddress
	}
	if len(req.DexIds) > 0 {
		params["dexIds"] = strings.Join(req.DexIds, ",")
	}
	if req.PriceImpactProtectionPercentage != "" {
		params["priceImpactProtectionPercentage"] = req.PriceImpactProtectionPercentage
	}
	if req.ComputeUnitPrice != "" {
		params["computeUnitPrice"] = req.ComputeUnitPrice
	}
	if req.ComputeUnitLimit != "" {
		params["computeUnitLimit"] = req.ComputeUnitLimit
	}

	if err = d.tr.Get(ctx, "/api/v5/dex/aggregator/swap-instruction", params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
