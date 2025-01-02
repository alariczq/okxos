package swap

import (
	"context"
	"strings"

	"github.com/imzhongqi/okxos/errcode"
)

type GetQuotesRequest struct {
	// ChainId is the chain ID (e.g., 1 for Ethereum. See Chain IDs)
	ChainId string `json:"chainId"`
	// Amount is the input amount of a token to be sold (set in minimal divisible units,
	// e.g., 1.00 USDT set as 1000000, 1.00 DAI set as 1000000000000000000), you could get the minimal divisible units from Tokenlist.
	Amount string `json:"amount"`
	// FromTokenAddress is the contract address of a token to be sold (e.g., 0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee)
	FromTokenAddress string `json:"fromTokenAddress"`
	// ToTokenAddress is the contract address of a token to be bought (e.g., 0xa892e1fef8b31acc44ce78e7db0a2dc610f92d00)
	ToTokenAddress string `json:"toTokenAddress"`
	// DexIds is DexId of the liquidity pool for limited quotes, multiple combinations separated by , (e.g.,1,50,180, see liquidity list for more)
	DexIds []string `json:"dexIds"`
	// PriceImpactProtectionPercentage is the percentage (between 0 - 1.0) of the price impact allowed.
	PriceImpactProtectionPercentage string `json:"priceImpactProtectionPercentage"`
	// FeePercent is the percentage of fromTokenAmount will be sent to the referrer's address,
	// the rest will be set as the input amount to be sold. min percentage：0
	FeePercent string `json:"feePercent"`
}

type DexProtocol struct {
	Percent string `json:"percent"`
	DexName string `json:"dexName"`
}

type TokenInfo struct {
	Decimal              string `json:"decimal"`
	IsHoneyPot           bool   `json:"isHoneyPot"`
	TaxRate              string `json:"taxRate"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenSymbol          string `json:"tokenSymbol"`
	TokenUnitPrice       string `json:"tokenUnitPrice"`
}

type SubRouter struct {
	DexProtocol []*DexProtocol `json:"dexProtocol"`
	FromToken   *TokenInfo     `json:"fromToken"`
	ToToken     *TokenInfo     `json:"toToken"`
}

type DexRouter struct {
	Router        string       `json:"router"`
	RouterPercent string       `json:"routerPercent"`
	SubRouterList []*SubRouter `json:"subRouterList"`
}

type QuoteCompare struct {
	AmountOut string `json:"amountOut"`
	DexLogo   string `json:"dexLogo"`
	DexName   string `json:"dexName"`
	TradeFee  string `json:"tradeFee"`
}

type QuotesResult struct {
	ChainId          string          `json:"chainId"`
	DexRouterList    []*DexRouter    `json:"dexRouterList"`
	EstimateGasFee   string          `json:"estimateGasFee"`
	FromToken        TokenInfo       `json:"fromToken"`
	FromTokenAmount  string          `json:"fromTokenAmount"`
	PriceImpactPct   string          `json:"priceImpactPct"`
	QuoteCompareList []*QuoteCompare `json:"quoteCompareList"`
	ToToken          TokenInfo       `json:"toToken"`
	ToTokenAmount    string          `json:"toTokenAmount"`
	TradeFee         string          `json:"tradeFee"`
}

// Get Quotes
func (s *SwapAPI) GetQuotes(ctx context.Context, req *GetQuotesRequest) (result *QuotesResult, err error) {
	params := map[string]string{
		"chainId":                         req.ChainId,
		"amount":                          req.Amount,
		"fromTokenAddress":                req.FromTokenAddress,
		"toTokenAddress":                  req.ToTokenAddress,
		"dexIds":                          strings.Join(req.DexIds, ","),
		"priceImpactProtectionPercentage": req.PriceImpactProtectionPercentage,
		"feePercent":                      req.FeePercent,
	}
	var results []*QuotesResult
	if err = s.tr.Get(ctx, "/api/v5/dex/aggregator/quote", params, &results); err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
