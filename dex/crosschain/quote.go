package crosschain

import (
	"context"
	"strconv"
	"strings"

	"github.com/imzhongqi/okxos/errcode"
)

// GetQuoteRequest Get quote request
// Docs: https://www.okx.com/web3/build/docs/waas/dex-get-route-information#get-route-information
type GetQuoteRequest struct {
	FromChainId                     string   `json:"fromChainId"`
	ToChainId                       string   `json:"toChainId"`
	FromTokenAddress                string   `json:"fromTokenAddress"`
	ToTokenAddress                  string   `json:"toTokenAddress"`
	Amount                          string   `json:"amount"`
	Slippage                        string   `json:"slippage"`
	Sort                            int      `json:"sort"`
	FeePercent                      string   `json:"feePercent"`
	AllowBridge                     []string `json:"allowBridge"`
	DenyBridge                      []string `json:"denyBridge"`
	PriceImpactProtectionPercentage string   `json:"priceImpactProtectionPercentage"`
}

type QuoteTokenInfo struct {
	Decimals             int64  `json:"decimals"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenSymbol          string `json:"tokenSymbol"`
}

type DexProtocol struct {
	Percent string `json:"percent"`
	DexName string `json:"dexName"`
}

type SubRouter struct {
	DexProtocol *DexProtocol    `json:"dexProtocol"`
	FromToken   *QuoteTokenInfo `json:"fromToken"`
	ToToken     *QuoteTokenInfo `json:"toToken"`
}

type DexRouter struct {
	Router        string       `json:"router"`
	RouterPercent string       `json:"routerPercent"`
	SubRouterList []*SubRouter `json:"subRouterList"`
}

type BridgeRouter struct {
	BridgeId                  int    `json:"bridgeId"`
	BridgeName                string `json:"bridgeName"`
	CrossChainFee             string `json:"crossChainFee"`
	CrossChainFeeTokenAddress string `json:"crossChainFeeTokenAddress"`
	OtherNativeFee            string `json:"otherNativeFee"`
	EstimateGasFee            string `json:"estimateGasFee"`
	EstimatedTime             string `json:"estimatedTime"`
}

type Router struct {
	EstimateTime        string        `json:"estimateTime"`
	EstimateGasFee      string        `json:"estimateGasFee"`
	FromChainNetworkFee string        `json:"fromChainNetworkFee"`
	ToChainNetworkFee   string        `json:"toChainNetworkFee"`
	ToTokenAmount       string        `json:"toTokenAmount"`
	MinimumReceived     string        `json:"minimumReceived"`
	NeedApprove         int           `json:"needApprove"`
	Router              *BridgeRouter `json:"router"`
	FromDexRouterList   []*DexRouter  `json:"fromDexRouterList"`
	ToDexRouterList     []*DexRouter  `json:"toDexRouterList"`
}

/*
{
  "code": "0",
  "data": [
    {
      "fromChainId": "501",
      "fromToken": {
        "decimals": 6,
        "tokenContractAddress": "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
        "tokenSymbol": "USDC"
      },
      "fromTokenAmount": "10000000000",
      "routerList": [
        {
          "estimateGasFee": "211587",
          "estimateTime": "430",
          "fromChainNetworkFee": "591000",
          "fromDexRouterList": [],
          "minimumReceived": "9996895509",
          "needApprove": 1,
          "router": {
            "bridgeId": 662,
            "bridgeName": "Circle-Bridge",
            "crossChainFee": "3.104491",
            "crossChainFeeTokenAddress": "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
            "otherNativeFee": "0"
          },
          "toChainNetworkFee": "0",
          "toDexRouterList": [],
          "toTokenAmount": "9996895509"
        }
      ],
      "toChainId": "1",
      "toToken": {
        "decimals": 6,
        "tokenContractAddress": "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
        "tokenSymbol": "USDC"
      }
    }
  ],
  "msg": ""
}
*/

type QuoteResult struct {
	FromChainId     string         `json:"fromChainId"`
	ToChainId       string         `json:"toChainId"`
	FromTokenAmount string         `json:"fromTokenAmount"`
	FromToken       QuoteTokenInfo `json:"fromToken"`
	ToToken         QuoteTokenInfo `json:"toToken"`
	RouterList      []*Router      `json:"routerList"`
}

// GetQuote Get quote
func (c *CrossChainAPI) GetQuote(ctx context.Context, quote *GetQuoteRequest) (result *QuoteResult, err error) {
	params := map[string]string{
		"fromChainId":      quote.FromChainId,
		"toChainId":        quote.ToChainId,
		"fromTokenAddress": quote.FromTokenAddress,
		"toTokenAddress":   quote.ToTokenAddress,
		"amount":           quote.Amount,
		"slippage":         quote.Slippage,
	}
	if quote.Sort != 0 {
		params["sort"] = strconv.Itoa(quote.Sort)
	}
	if quote.FeePercent != "" {
		params["feePercent"] = quote.FeePercent
	}
	if len(quote.AllowBridge) > 0 {
		params["allowBridge"] = strings.Join(quote.AllowBridge, ",")
	}
	if len(quote.DenyBridge) > 0 {
		params["denyBridge"] = strings.Join(quote.DenyBridge, ",")
	}
	if quote.PriceImpactProtectionPercentage != "" {
		params["priceImpactProtectionPercentage"] = quote.PriceImpactProtectionPercentage
	}

	var results []*QuoteResult
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/quote", params, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
