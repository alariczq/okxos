package dex

import (
	"github.com/imzhongqi/okxos/client"
	"github.com/imzhongqi/okxos/dex/crosschain"
)

type DexAPI struct {
	CrossChain *crosschain.CrossChainAPI
	tr         client.Transport
}

func NewDexAPI(tr client.Transport) *DexAPI {
	return &DexAPI{
		tr:         tr,
		CrossChain: crosschain.NewCrossChainAPI(tr),
	}
}
