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
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/imzhongqi/okxos/errcode"
)

type ValidateAddressRequest struct {
	// Unique identifier for the chain
	ChainIndex string `json:"chainIndex"`
	// Address
	Address string `json:"address"`
}

// AddressType
// 0: Invalid address format
// 1: Valid user address
// 2: Valid contract address
type AddressType string

func (a AddressType) IsInvalidAddress() bool {
	return a == "0"
}

func (a AddressType) IsUserAddress() bool {
	return a == "1"
}

func (a AddressType) IsContractAddress() bool {
	return a == "2"
}

type ValidateAddressResult struct {
	AddressType AddressType `json:"addressType"`
	// Whether the address has hit blacklist check
	HitBlacklist bool `json:"hitBlacklist"`
	// Tag type for blacklisted addresses, including phishing, contract vulnerabilities, etc.
	Tag string `json:"tag"`
}

// ValidateAddress Provide an address to determine if it is a valid user or contract address, and whether it has hit blacklist check.
func (w *WalletAPI) ValidateAddress(ctx context.Context, req *ValidateAddressRequest) (result *ValidateAddressResult, err error) {
	params := map[string]string{
		"chainIndex": req.ChainIndex,
		"address":    req.Address,
	}
	var results []*ValidateAddressResult
	err = w.tr.Get(ctx, "/api/v5/wallet/pre-transaction/validate-address", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type TransactionBroadcastRequest struct {
	SignedTx   string `json:"signedTx"`
	ChainIndex string `json:"chainIndex"`
	Address    string `json:"address"`
	AccountId  string `json:"accountId"`
}

type TransactionBroadcastResult struct {
	OrderId string `json:"orderId"`
}

func (w *WalletAPI) TransactionBroadcast(ctx context.Context, tx *TransactionBroadcastRequest) (result *TransactionBroadcastResult, err error) {
	var results []*TransactionBroadcastResult
	err = w.tr.Post(ctx, "/api/v5/wallet/pre-transaction/broadcast-transaction", tx, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetNonceRequest struct {
	// Unique identifier for the chain
	ChainIndex string `json:"chainIndex"`
	// Address
	Address string `json:"address"`
}

type GetNonceResult struct {
	Nonce        string `json:"nonce"`
	PendingNonce string `json:"pendingNonce"`
}

func (w *WalletAPI) GetNonce(ctx context.Context, req *GetNonceRequest) (result *GetNonceResult, err error) {
	params := map[string]string{
		"chainIndex": req.ChainIndex,
		"address":    req.Address,
	}
	var results []*GetNonceResult
	err = w.tr.Get(ctx, "/api/v5/wallet/pre-transaction/nonce", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetSuiObjectRequest struct {
	// Unique identifier for the chain
	ChainIndex string `json:"chainIndex"`
	// Wallet address
	Address string `json:"address"`
	// Token address
	TokenAddress string `json:"tokenAddress"`
	// Number of entries per query, default is 50, maximum is 50
	Limit string `json:"limit"`
	// Cursor position, defaults to the first entry
	Cursor string `json:"cursor"`
}

type SuiObject struct {
	// Amount Token balance
	Amount string `json:"amount"`
	// Digest 32-byte transaction summary indicating the last transaction that included this object as an output
	Digest string `json:"digest"`
	// Version 8-byte unsigned integer version, monotonically increasing with each transaction that modifies it
	Version string `json:"version"`
	// A 32-byte globally unique ID. The object ID is derived from the digest of the transaction that created the object and a counter of the number of IDs generated by the encoding transaction.
	ObjectId string `json:"objectId"`
}

type SuiObjectResult struct {
	// Token address
	TokenAddress string `json:"tokenAddress"`
	// Next cursor
	Cursor string `json:"cursor"`
	// Object list
	Objects []*SuiObject `json:"objects"`
}

func (w *WalletAPI) GetSuiObject(ctx context.Context, req *GetSuiObjectRequest) (result *SuiObjectResult, err error) {
	params := map[string]string{
		"chainIndex":   req.ChainIndex,
		"address":      req.Address,
		"tokenAddress": req.TokenAddress,
	}
	if req.Limit != "" {
		params["limit"] = req.Limit
	}
	if req.Cursor != "" {
		params["cursor"] = req.Cursor
	}
	var results []*SuiObjectResult
	err = w.tr.Get(ctx, "/api/v5/wallet/pre-transaction/sui-object", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

type GetSignInfoRequest struct {
	// Unique identifier for the chain
	ChainIndex string `json:"chainIndex"`
	// From address
	FromAddr string `json:"fromAddr"`
	// To address
	ToAddr string `json:"toAddr"`
	// Native token amount for the transaction, default is 0.
	// Must include this parameter when involving mainnet coin transfer,
	// otherwise it will affect the calculation of gas limit,
	// passed in the smallest unit of the chain native token, such as ETH wei.
	TxAmount string `json:"txAmount"`
	// Extension parameters, used to add calldata and other information
	ExtJson *ExtJson `json:"extJson,omitempty"`
}

type ExtJson struct {
	// Only appliable to EVM
	InputData string `json:"inputData"`
	// Only appliable to EVM
	// 1: BRC-20
	// 2: ARC-20
	// 3: Runes
	// 4: ordi_nft
	Protocol string `json:"protocol"`
	// Only appliable to Solana
	TokenAddress string `json:"tokenAddress"`
	// Only appliable to Tron.
	// 1:owner permission
	// 2: witness permission
	// 1 by default
	PermissionType string `json:"permissionType"`
	// Only appliable to Tron. Required if interact with contract, 30000000 by default
	FeeLimit string `json:"feeLimit"`
}

type SignInfoEvm struct {
	GasLimit string    `json:"gasLimit"`
	Nonce    string    `json:"nonce"`
	GasPrice *GasPrice `json:"gasPrice"`
}

type Eip1559Protocol struct {
	BaseFee            string `json:"baseFee"`
	FastPriorityFee    string `json:"fastPriorityFee"`
	SafePriorityFee    string `json:"safePriorityFee"`
	SuggestGasPrice    string `json:"suggestGasPrice"`
	ProposePriorityFee string `json:"proposePriorityFee"`
}

type GasPrice struct {
	Normal           string           `json:"normal"`
	Min              string           `json:"min"`
	Max              string           `json:"max"`
	SupportedEip1559 bool             `json:"supportedEip1559"`
	Eip1559Protocol  *Eip1559Protocol `json:"eip1559Protocol"`
}

type SignInfoUtxo struct {
	NormalFeeRate     string `json:"normalFeeRate"`
	MaxFeeRate        string `json:"maxFeeRate"`
	MinFeeRate        string `json:"minFeeRate"`
	InscriptionOutput string `json:"inscriptionOutput"`
	MinOutput         string `json:"minOutput"`
	NormalCost        string `json:"normalCost"`
	MaxCost           string `json:"maxCost"`
	MinCost           string `json:"minCost"`
}

type SignInfoSolana struct {
	BaseFee              string            `json:"baseFee"`
	PriorityFee          *PriorityFee      `json:"priorityFee"`
	RecentBlockHash      string            `json:"recentBlockHash"`
	LastValidBlockHeight string            `json:"lastValidBlockHeight"`
	FromAddressRent      string            `json:"fromAddressRent"`
	ToAddressRent        string            `json:"toAddressRent"`
	TokenAccountInfo     *TokenAccountInfo `json:"tokenAccountInfo"`
}

type PriorityFee struct {
	// Normal unit price for the transaction
	NormalUnitPrice string `json:"normalUnitPrice"`
	// Min unit price for the transaction
	MinUnitPrice string `json:"minUnitPrice"`
	// Max unit price for the transaction
	MaxUnitPrice string `json:"maxUnitPrice"`
}

type TokenAccountInfo struct {
	Lamports string `json:"lamports"`
	// OwnerAddress from address
	OwnerAddress string `json:"ownerAddress"`
	// MintAddress token address
	MintAddress         string `json:"mintAddress"`
	TokenAccountAddress string `json:"tokenAccountAddress"`
	Decimal             string `json:"decimal"`
}

type SignInfoTron struct {
	Fee string `json:"fee"`
	// RefBlockBytes Reference block bytes The 6th to 8th bytes (not included) of the reference block height are used to help
	// verify whether the transaction is based on the valid state of the current blockchain and prevent forked transaction replay
	RefBlockBytes string `json:"refBlockBytes"`
	// RefBlockHash Reference block hash The 8th to 16th bytes (not included) of the reference block hash are used.
	// If the transaction's ref_block_hash does not match the actual existing block hash,
	// the transaction may be rejected or marked as invalid.
	RefBlockHash string `json:"refBlockHash"`
	// Expiration for the transaction
	Expiration string `json:"expiration"`
	// Timestamp for the transaction
	Timestamp string `json:"timestamp"`
}

type SignInfoResult struct {
	Evm    *SignInfoEvm
	Utxo   *SignInfoUtxo
	Solana *SignInfoSolana
	Tron   *SignInfoTron
}

func (s *SignInfoResult) UnmarshalJSON(data []byte) (err error) {
	r := bytes.NewReader(data)
	decode := func(v any) error {
		r.Seek(0, io.SeekStart)
		decoder := json.NewDecoder(r)
		decoder.DisallowUnknownFields()
		decoder.More()

		return decoder.Decode(v)
	}

	evm := new(SignInfoEvm)
	if err = decode(evm); err == nil {
		s.Evm = evm
		return nil
	}

	utxo := new(SignInfoUtxo)
	if err = decode(utxo); err == nil {
		s.Utxo = utxo
		return nil
	}

	solana := new(SignInfoSolana)
	if err = decode(solana); err == nil {
		s.Solana = solana
		return nil
	}

	tron := new(SignInfoTron)
	if err = decode(tron); err == nil {
		s.Tron = tron
		return nil
	}

	return err
}

func (w *WalletAPI) GetSignInfo(ctx context.Context, req *GetSignInfoRequest) (result *SignInfoResult, err error) {
	var results []*SignInfoResult
	err = w.tr.Post(ctx, "/api/v5/wallet/pre-transaction/sign-info", req, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
