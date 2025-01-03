package crosschain

import (
	"context"

	"github.com/imzhongqi/okxos/errcode"
)

type CrossChainFee struct {
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type CrossChainInfo struct {
	Memo                   string `json:"memo"`
	DestinationChainGasfee string `json:"destinationChainGasfee"`
	DetailStatus           string `json:"detailStatus"`
	Status                 string `json:"status"`
}

type TransactionStatus struct {
	FromChainId        string         `json:"fromChainId"`
	ToChainId          string         `json:"toChainId"`
	FromTxHash         string         `json:"fromTxHash"`
	ToTxHash           string         `json:"toTxHash"`
	FromAmount         string         `json:"fromAmount"`
	FromTokenAddress   string         `json:"fromTokenAddress"`
	ToAmount           string         `json:"toAmount"`
	ToTokenAddress     string         `json:"toTokenAddress"`
	ErrorMsg           string         `json:"errorMsg"`
	BridgeHash         string         `json:"bridgeHash"`
	RefundChainId      string         `json:"refundChainId"`
	RefundTokenAddress string         `json:"refundTokenAddress"`
	RefundTxHash       string         `json:"refundTxHash"`
	SourceChainGasfee  string         `json:"sourceChainGasfee"`
	CrossChainFee      CrossChainFee  `json:"crossChainFee"`
	Symbol             string         `json:"symbol"`
	Address            string         `json:"address"`
	Amount             string         `json:"amount"`
	CrossChainInfo     CrossChainInfo `json:"crossChainInfo"`
}

type GetTransactionStatusRequest struct {
	// Hash address of the source chain
	Hash string `json:"hash"`
	// ChainId Source chain ID (e.g., 1 for Ethereum. See Chain IDs)
	ChainId string `json:"chainId"`
}

// GetTransactionStatus Check the final status of the cross-chain swap according to transaction hash.
func (c *CrossChainAPI) GetTransactionStatus(ctx context.Context, req *GetTransactionStatusRequest) (result *TransactionStatus, err error) {
	params := map[string]string{
		"hash": req.Hash,
	}
	if req.ChainId != "" {
		params["chainId"] = req.ChainId
	}
	var results []*TransactionStatus
	if err = c.tr.Get(ctx, "/api/v5/dex/cross-chain/status", params, &results); err != nil {
		return
	}
	if len(results) == 0 {
		return nil, errcode.ErrResultsNotFound
	}
	return results[0], nil
}
