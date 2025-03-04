package crosschain

import (
	"context"
)

type ChainInfo struct {
	ChainId                string `json:"chainId"`
	ChainName              string `json:"chainName"`
	DexTokenApproveAddress string `json:"dexTokenApproveAddress"`
}

// GetSupportedChains Get supported chains
func (c *CrossChainAPI) GetSupportedChains(ctx context.Context, chainId string) (result []ChainInfo, err error) {
	params := map[string]string{}
	if chainId != "" {
		params["chainId"] = chainId
	}
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/supported/chain", params, &result)
	return
}
