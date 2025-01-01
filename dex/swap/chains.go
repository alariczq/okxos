package swap

import (
	"context"
	"strconv"
)

type ChainInfo struct {
	ChainId             int64  `json:"chainId"`
	ChainName           string `json:"chainName"`
	DexTokenApproveAddr string `json:"dexTokenApproveAddress"`
}

// Get Supported Chains
func (d *SwapAPI) GetSupportedChains(ctx context.Context, chainId ...int64) (result []*ChainInfo, err error) {
	params := map[string]string{}
	if len(chainId) > 0 && chainId[0] != 0 {
		params["chainId"] = strconv.FormatInt(chainId[0], 10)
	}
	err = d.tr.Get(ctx, "/api/v5/dex/aggregator/supported/chain", params, &result)
	return
}
