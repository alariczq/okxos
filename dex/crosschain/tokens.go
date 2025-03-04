package crosschain

import (
	"context"
)

type TokenInfo struct {
	ChainId              string `json:"chainId"`
	Decimals             int    `json:"decimals"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenLogoUrl         string `json:"tokenLogoUrl"`
	TokenName            string `json:"tokenName"`
	TokenSymbol          string `json:"tokenSymbol"`
}

// GetSupportedTokens List of tokens available for traded directly across the cross-chain bridge.
func (c *CrossChainAPI) GetSupportedTokens(ctx context.Context, chainId string) (result []TokenInfo, err error) {
	params := map[string]string{}
	if chainId != "" {
		params["chainId"] = chainId
	}
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/supported/tokens", params, &result)
	return
}

type TokenPair struct {
	FromChainId      string `json:"fromChainId"`
	ToChainId        string `json:"toChainId"`
	FromTokenAddress string `json:"fromTokenAddress"`
	ToTokenAddress   string `json:"toTokenAddress"`
	FromTokenSymbol  string `json:"fromTokenSymbol"`
	ToTokenSymbol    string `json:"toTokenSymbol"`
}

// GetSupportedBridgeTokensPairs List of tokens pairs available for traded directly across the cross-chain bridge.
func (c *CrossChainAPI) GetSupportedBridgeTokensPairs(ctx context.Context, fromChainId string) (result []TokenPair, err error) {
	params := map[string]string{
		"fromChainId": fromChainId,
	}
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/supported/bridge-tokens-pairs", params, &result)
	return
}
