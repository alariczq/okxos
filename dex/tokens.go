package dex

import (
	"context"
	"strconv"
)

type Tokens struct {
	Decimals             string `json:"decimals"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenLogoUrl         string `json:"tokenLogoUrl"`
	TokenName            string `json:"tokenName"`
	TokenSymbol          string `json:"tokenSymbol"`
}

// Get tokens
// It fetches a list of tokens. This interface returns a list of tokens that belong to major platforms or
// are deemed significant enough by OKX. However, you can still quote and swap other tokens outside of this list on OKX DEX.
func (d *DexAPI) GetSupportedTokens(ctx context.Context, chainId int64) (result []*Tokens, err error) {
	params := map[string]string{
		"chainId": strconv.FormatInt(chainId, 10),
	}
	err = d.tr.Get(ctx, "/api/v5/dex/aggregator/all-tokens", params, &result)
	return
}
