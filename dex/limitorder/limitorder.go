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

package limitorder

import (
	"context"

	"github.com/imzhongqi/okxos/client"
	"github.com/imzhongqi/okxos/errcode"
)

// LimitOrderAPI provides API operations for OKX Web3 limit orders
type LimitOrderAPI struct {
	tr client.Transport
}

// NewLimitOrderAPI creates a new LimitOrderAPI instance
func NewLimitOrderAPI(tr client.Transport) *LimitOrderAPI {
	return &LimitOrderAPI{
		tr: tr,
	}
}

// OrderData represents the limit order data structure
type OrderData struct {
	Salt          string `json:"salt"`
	MakerToken    string `json:"makerToken"`
	TakerToken    string `json:"takerToken"`
	Maker         string `json:"maker"`
	Receiver      string `json:"receiver"`
	AllowedSender string `json:"allowedSender"`
	MakingAmount  string `json:"makingAmount"`
	TakingAmount  string `json:"takingAmount"`
	MinReturn     string `json:"minReturn"`
	DeadLine      string `json:"deadLine"`
	PartiallyAble bool   `json:"partiallyAble"`
}

// CreateOrderRequest represents the request parameters for creating a limit order
type CreateOrderRequest struct {
	OrderHash string    `json:"orderHash"`
	ChainId   string    `json:"chainId"`
	Signature string    `json:"signature"`
	Data      OrderData `json:"data"`
}

// OrderDetail represents the response data for a limit order
type OrderDetail struct {
	ChainId              string `json:"chainId"`
	CreateTime           string `json:"createTime"`
	ExpireTime           string `json:"expireTime"`
	MakerAssetAddress    string `json:"makerAssetAddress"`
	MakerRate            string `json:"makerRate"`
	MakerTokenAddress    string `json:"makerTokenAddress"`
	MakingAmount         string `json:"makingAmount"`
	OrderHash            string `json:"orderHash"`
	Receiver             string `json:"receiver"`
	RemainingMakerAmount string `json:"remainingMakerAmount"`
	Salt                 string `json:"salt"`
	Signature            string `json:"signature"`
	Status               string `json:"status"`
	TakerAssetAddress    string `json:"takerAssetAddress"`
	TakerRate            string `json:"takerRate"`
	TakerTokenAddress    string `json:"takerTokenAddress"`
	TakingAmount         string `json:"takingAmount"`
}

// CreateOrder creates a limit order
func (api *LimitOrderAPI) CreateOrder(ctx context.Context, req CreateOrderRequest) (*OrderDetail, error) {
	var result []OrderDetail
	err := api.tr.Post(ctx, "/dex/aggregator/limit-order/save-order", req, &result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errcode.ErrResultsNotFound
	}

	return &result[0], nil
}

// ListOrdersRequest represents the request parameters for listing limit orders
type ListOrdersRequest struct {
	ChainId    string `json:"chainId"`
	Page       string `json:"page,omitempty"`
	Limit      string `json:"limit,omitempty"`
	Statuses   string `json:"statuses,omitempty"`
	TakerAsset string `json:"takerAsset,omitempty"`
	MakerAsset string `json:"makerAsset,omitempty"`
}

// ListOrders lists limit orders based on the provided parameters
func (api *LimitOrderAPI) ListOrders(ctx context.Context, req ListOrdersRequest) ([]*OrderDetail, error) {
	params := map[string]string{
		"chainId": req.ChainId,
	}

	if req.Page != "" {
		params["page"] = req.Page
	}

	if req.Limit != "" {
		params["limit"] = req.Limit
	}

	if req.Statuses != "" {
		params["statuses"] = req.Statuses
	}

	if req.TakerAsset != "" {
		params["takerAsset"] = req.TakerAsset
	}

	if req.MakerAsset != "" {
		params["makerAsset"] = req.MakerAsset
	}

	var result []*OrderDetail
	err := api.tr.Get(ctx, "/dex/aggregator/limit-order/all", params, &result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errcode.ErrResultsNotFound
	}

	return result, nil
}

// GetOrderRequest represents the request parameters for getting a limit order
type GetOrderRequest struct {
	ChainId   string `json:"chainId"`
	OrderHash string `json:"orderHash"`
}

// GetOrder gets the details of a specific limit order
func (api *LimitOrderAPI) GetOrder(ctx context.Context, chainId, orderHash string) (*OrderDetail, error) {
	params := map[string]string{
		"chainId":   chainId,
		"orderHash": orderHash,
	}

	var result []OrderDetail
	err := api.tr.Get(ctx, "/dex/aggregator/limit-order/detail", params, &result)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

// CancelOrderRequest represents the request parameters for canceling a limit order
type CancelOrderRequest struct {
	OrderHash string `json:"orderHash"`
}

// CancelOrder gets the calldata for canceling a limit order
func (api *LimitOrderAPI) CancelOrder(ctx context.Context, chainId, orderHash string) (calldata string, err error) {
	params := map[string]string{
		"orderHash": orderHash,
	}

	err = api.tr.Get(ctx, "/dex/aggregator/limit-order/cancel/calldata", params, &calldata)
	if err != nil {
		return "", err
	}

	return calldata, nil
}
