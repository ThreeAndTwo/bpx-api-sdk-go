package constants

import "errors"

var (
	ErrURLIsNull       = errors.New("url is null")
	ErrApiKeyIsNull    = errors.New("apiKey is null")
	ErrApiSecretIsNull = errors.New("apiSecret is null")
	ErrSymbolIsNull    = errors.New("symbol is null")
)
