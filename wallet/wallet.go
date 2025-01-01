package wallet

import (
	"github.com/imzhongqi/okxos/client"
)

type WalletAPI struct {
	tr client.Transport
}

func NewWalletAPI(client client.Transport) *WalletAPI {
	return &WalletAPI{
		tr: client,
	}
}
