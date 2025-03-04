package dex

import (
	"context"
	"strconv"
)

type Liquidity struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

// Get Liquidity
func (d *DexAPI) GetLiquidity(ctx context.Context, chainId int64) (result []Liquidity, err error) {
	params := map[string]string{
		"chainId": strconv.FormatInt(chainId, 10),
	}
	err = d.tr.Get(ctx, "/api/v5/dex/aggregator/get-liquidity", params, &result)
	return
}
