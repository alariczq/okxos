package wallet

import (
	"context"
)

type SubscribeRequest struct {
	Url string `json:"url"`
	// Type of data to subscribe to. Valid values are:
	// block: triggers when any block is generated
	// token_issuance: Token issuance
	// fee_fluctuation: Fee fluctuation
	Type string `json:"type"`
	// Chain Index
	ChainIndex string `json:"chainIndex"`
	// Name of the subscription
	Name string `json:"name"`
	// Fee fluctuation filter, applicable only when the type is fee_fluctuation
	FeeChangeFilter *FeeChangeFilter `json:"feeChangeFilter,omitempty"`
}

type FeeChangeFilter struct {
	// Minimum fluctuation
	MinChange string `json:"minChange"`
	// Maximum fluctuation
	MaxChange string `json:"maxChange"`
}

type SubscribeResult struct {
	Id string `json:"id"`
}

// Subscribe Webhook
func (w *WalletAPI) Subscribe(ctx context.Context, req []*SubscribeRequest) (results []*SubscribeResult, err error) {
	err = w.tr.Post(ctx, "/api/v5/wallet/webhook/subscribe", req, &results)
	return
}

type UnsubscribeRequest struct {
	Id string `json:"id"`
}

type UnsubscribeResult struct {
	Id string `json:"id"`
}

// Unsubscribe Webhook
func (w *WalletAPI) Unsubscribe(ctx context.Context, reqs []*UnsubscribeRequest) (results []*UnsubscribeResult, err error) {
	err = w.tr.Post(ctx, "/api/v5/wallet/webhook/unsubscribe", reqs, &results)
	return
}

type SubscriptionListResult struct {
	Id         string `json:"id"`
	ChainIndex string `json:"chainIndex"`
	Name       string `json:"name"`
	Url        string `json:"url"`
	Type       string `json:"type"`
}

// Subscription List
func (w *WalletAPI) SubscriptionList(ctx context.Context) (results []*SubscriptionListResult, err error) {
	err = w.tr.Get(ctx, "/api/v5/wallet/webhook/subscriptions", nil, &results)
	return
}
