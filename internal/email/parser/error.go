package parser

import "errors"

var (
	ErrParserNotSupported = errors.New("parser not supported")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrAmountNotFound     = errors.New("amount not found")
	ErrMerchantNotFound   = errors.New("merchant not found")
)
