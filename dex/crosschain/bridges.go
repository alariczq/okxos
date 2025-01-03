package crosschain

import "context"

/*
Get supported bridges#
Get information of the cross-chain bridges supported by OKX’s DEX cross-chain aggregator protocol.

Request address#
GET https://www.okx.com/api/v5/dex/cross-chain/supported/bridges

Request param#
Parameter	Type	Required	Description
chainId	String	No	Chain ID (e.g., 1 for Ethereum. See Chain IDs)
Response param#
Parameter	Type	Required	Description
bridgeName	String	Yes	Name of bridge (e.g., cBridge)
bridgeId	Integer	Yes	Bridge ID (e.g., 211)
requiredOtherNativeFee	boolean	Yes	if this bridge require native fee
logo	String	Yes	Bridge Logo URL
supportedChains	Array	Yes	Return chain id that bridge supported
*/

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
