package accounts

import "errors"

var (
	errDoRequest          = errors.New("error doing request")
	errResponseReadBody   = errors.New("error reading response body")
	errResponseStatusCode = errors.New("response error")
)
