package wallet

import "github.com/imzhongqi/okxos/errcode"

// IsBlockchainNotSupported 81104 Blockchain not supported
func IsBlockchainNotSupported(err error) bool {
	return errcode.Is(err, 81104)
}

// IsWalletVerificationError 81105 Wallet verification error
func IsWalletVerificationError(err error) bool {
	return errcode.Is(err, 81105)
}

// IsAddressMustBeLowercase 81106 Address must be in lowercase
func IsAddressMustBeLowercase(err error) bool {
	return errcode.Is(err, 81106)
}

// IsTooManyWalletAddresses 81107 Too many wallet addresses
func IsTooManyWalletAddresses(err error) bool {
	return errcode.Is(err, 81107)
}

// IsWalletTypeMismatch 81108 Wallet type mismatch
func IsWalletTypeMismatch(err error) bool {
	return errcode.Is(err, 81108)
}

// IsAddressUpdateError 81109 Address update error
func IsAddressUpdateError(err error) bool {
	return errcode.Is(err, 81109)
}

// IsChainNotSupported 81150 Chain not supported in this interface
func IsChainNotSupported(err error) bool {
	return errcode.Is(err, 81150)
}

// IsTokenAddressIncorrect 81151 Token address incorrect
func IsTokenAddressIncorrect(err error) bool {
	return errcode.Is(err, 81151)
}

// IsTokenDoesNotExist 81152 Token does not exist
func IsTokenDoesNotExist(err error) bool {
	return errcode.Is(err, 81152)
}

// IsTokenIsPlatformToken 81153 This token is a platform token, no need to add
func IsTokenIsPlatformToken(err error) bool {
	return errcode.Is(err, 81153)
}

// IsBlockchainAndAddressDoNotMatch 81157 Blockchain and address do not match
func IsBlockchainAndAddressDoNotMatch(err error) bool {
	return errcode.Is(err, 81157)
}

// IsTokenProtocolNotSupported 81158 Token protocol not supported
func IsTokenProtocolNotSupported(err error) bool {
	return errcode.Is(err, 81158)
}

// IsDataCaching 81159 Data caching, please try again later
func IsDataCaching(err error) bool {
	return errcode.Is(err, 81159)
}

// IsTransactionNotFound 81201 Transaction not found
func IsTransactionNotFound(err error) bool {
	return errcode.Is(err, 81201)
}

// IsTransactionStillPending 81202 Transaction still pending
func IsTransactionStillPending(err error) bool {
	return errcode.Is(err, 81202)
}

// IsExtjsonParametersNotFound 81203 Transaction extjson parameters not found
func IsExtjsonParametersNotFound(err error) bool {
	return errcode.Is(err, 81203)
}

// IsFromAddressMismatchAccount 81302 FromAddress does not belong to the account ID
func IsFromAddressMismatchAccount(err error) bool {
	return errcode.Is(err, 81302)
}

// IsInsufficientBalanceToPay 81351 Insufficient balance to pay
func IsInsufficientBalanceToPay(err error) bool {
	return errcode.Is(err, 81351)
}

// IsAddressIsIllegal 81353 Address is illegal
func IsAddressIsIllegal(err error) bool {
	return errcode.Is(err, 81353)
}

// IsNodeReturnFailed 81451 Node return failed
func IsNodeReturnFailed(err error) bool {
	return errcode.Is(err, 81451)
}
