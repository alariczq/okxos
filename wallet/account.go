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

	"github.com/imzhongqi/okxos/errcode"
)

type Address struct {
	ChainIndex string `json:"chainIndex"`
	Address    string `json:"address"`
}

type CreateAccountRequest struct {
	Addresses []*Address `json:"addresses"`
}

type CreateAccountResult struct {
	AccountId string `json:"accountId"`
}

func (w *WalletAPI) CreateAccount(ctx context.Context, req *CreateAccountRequest) (result *CreateAccountResult, err error) {
	var results []*CreateAccountResult
	err = w.tr.Post(ctx, "/api/v5/wallet/account/create-wallet-account", req, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}

func (w *WalletAPI) DeleteAccount(ctx context.Context, accountId string) (err error) {
	err = w.tr.Post(ctx, "/api/v5/wallet/account/delete-account", map[string]string{"accountId": accountId}, nil)
	return
}

type UpdateType string

const (
	UpdateTypeAdd    UpdateType = "add"
	UpdateTypeDelete UpdateType = "delete"
)

type UpdateAccountRequest struct {
	AccountId  string     `json:"accountId"`
	UpdateType UpdateType `json:"updateType"`
	Addresses  []*Address `json:"addresses"`
}

func (w *WalletAPI) UpdateAccount(ctx context.Context, req *UpdateAccountRequest) (err error) {
	err = w.tr.Post(ctx, "/api/v5/wallet/account/update-wallet-account", req, nil)
	return
}

type Account struct {
	AccountId   string `json:"accountId"`
	AccountType string `json:"accountType"`
}

type GetAccountResult struct {
	Accounts []*Account `json:"accounts"`
	Cursor   string     `json:"cursor"`
}

func (w *WalletAPI) GetAccount(ctx context.Context, limit string, cursor ...string) (result *GetAccountResult, err error) {
	params := map[string]string{
		"limit": limit,
	}
	if len(cursor) > 0 && cursor[0] != "" {
		params["cursor"] = cursor[0]
	}
	var results []*GetAccountResult
	err = w.tr.Get(ctx, "/api/v5/wallet/account/accounts", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
