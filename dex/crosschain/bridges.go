package crosschain

import "context"

type BridgeInfo struct {
	BridgeName             string   `json:"bridgeName"`
	BridgeId               int64    `json:"bridgeId"`
	RequiredOtherNativeFee bool     `json:"requiredOtherNativeFee"`
	Logo                   string   `json:"logo"`
	SupportedChains        []string `json:"supportedChains"`
}

// GetSupportedBridges Get supported bridges
func (c *CrossChainAPI) GetSupportedBridges(ctx context.Context, chainId string) (result []*BridgeInfo, err error) {
	params := map[string]string{}
	if chainId != "" {
		params["chainId"] = chainId
	}
	err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/supported/bridges", params, &result)
	return
}
