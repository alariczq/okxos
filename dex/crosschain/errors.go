// Copyright (c) 2024-NOW imzhongqi <imzhongqi@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
