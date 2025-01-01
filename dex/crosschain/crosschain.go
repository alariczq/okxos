package crosschain

import "github.com/imzhongqi/okxos/client"

type CrossChainAPI struct {
	tr client.Transport
}

func NewCrossChainAPI(tr client.Transport) *CrossChainAPI {
	return &CrossChainAPI{
		tr: tr,
	}
}
