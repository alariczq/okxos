package crosschain

import "github.com/imzhongqi/okxos/errcode"

// IsInsufficientLiquidity 82000 Insufficient liquidity
func IsInsufficientLiquidity(err error) bool {
	return errcode.Is(err, 82000)
}

// IsCommissionServiceNotAvailable 82001 The commission service is not available during the upgrade
func IsCommissionServiceNotAvailable(err error) bool {
	return errcode.Is(err, 82001)
}

// IsMinimumAmount 82102 Minimum amount is {0}
func IsMinimumAmount(err error) bool {
	return errcode.Is(err, 82102)
}

// IsMaximumAmount 82103 Maximum amount is {0}
func IsMaximumAmount(err error) bool {
	return errcode.Is(err, 82103)
}

// IsThisTokenIsNotSupported 82104 This token is not supported
func IsThisTokenIsNotSupported(err error) bool {
	return errcode.Is(err, 82104)
}

// IsThisChainIsNotSupported 82105 This chain is not supported
func IsThisChainIsNotSupported(err error) bool {
	return errcode.Is(err, 82105)
}

// IsValueDifference 82112 The value difference from this transaction’s quote route is higher than {num}, which may cause asset loss.
func IsValueDifference(err error) bool {
	return errcode.Is(err, 82112)
}

// IsSlippageTooLow 82114 The slippage too low,Suggest {0}
func IsSlippageTooLow(err error) bool {
	return errcode.Is(err, 82114)
}

// IsChainHasNotTokenPairs 82115 The chain has not token pairs
func IsChainHasNotTokenPairs(err error) bool {
	return errcode.Is(err, 82115)
}

// IsCrossChainBridgeNotFound 82116 No suitable cross-chain bridge found
func IsCrossChainBridgeNotFound(err error) bool {
	return errcode.Is(err, 82116)
}
