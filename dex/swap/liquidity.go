package swap

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
func (s *SwapAPI) GetLiquidity(ctx context.Context, chainId int64) (result []*Liquidity, err error) {
	params := map[string]string{
		"chainId": strconv.FormatInt(chainId, 10),
	}
	err = s.tr.Get(ctx, "/api/v5/dex/aggregator/get-liquidity", params, &result)
	return
}
