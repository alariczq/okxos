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

package dex

import "github.com/imzhongqi/okxos/errcode"

// IsRepeatedRequest 80000 Repeated request
func IsRepeatedRequest(err error) bool {
	return errcode.Is(err, 80000)
}

// IsCallDataExceedsMaxLimit 80001 CallData exceeds the maximum limit. Try again in 5 minutes.
func IsCallDataExceedsMaxLimit(err error) bool {
	return errcode.Is(err, 80001)
}

// IsTokenLimitReached 80002 Requested token Object count has reached the limit.
func IsTokenLimitReached(err error) bool {
	return errcode.Is(err, 80002)
}

// IsNativeTokenLimitReached 80003 Requested native token Object count has reached the limit.
func IsNativeTokenLimitReached(err error) bool {
	return errcode.Is(err, 80003)
}

// IsTimeoutQueryingSuiObject 80004 Timeout when querying SUI Object.
func IsTimeoutQueryingSuiObject(err error) bool {
	return errcode.Is(err, 80004)
}

// IsSuiObjectsNotEnough 82000 Not enough Sui objects under the address for swapping
func IsSuiObjectsNotEnough(err error) bool {
	return errcode.Is(err, 82000)
}

// IsInsufficientLiquidity 82001 Insufficient liquidity
func IsInsufficientLiquidity(err error) bool {
	return errcode.Is(err, 82001)
}

// IsValueDifference 82112 The value difference from this transaction’s quote route is higher than {num},
// which may cause asset loss,The default value is 90%.
// It can be adjusted using the string priceImpactProtectionPercentage.
func IsValueDifference(err error) bool {
	return errcode.Is(err, 82112)
}

// IsCallDataExceedsMaxLimit 82116 callData exceeds the maximum limit. Try again in 5 minutes.
// func IsCallDataExceedsMaxLimit(err error) bool {
// 	return errcode.Is(err, 82116)
// }

// IsTransactionIntercepted 82120 Detected honeypot tokens or high-risk tokens with a 100% buy/sell tax.
// Transactions have been intercepted.
func IsTransactionIntercepted(err error) bool {
	return errcode.Is(err, 82120)
}
