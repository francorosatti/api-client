package accounts

import "errors"

var (
	ErrAccountBadRequest = errors.New("account bad request")
	ErrAccountNotFound   = errors.New("account not found")
	ErrAccountConflict   = errors.New("account conflict with version")

	errDoRequest          = errors.New("error doing request")
	errResponseReadBody   = errors.New("error reading response body")
	errResponseStatusCode = errors.New("response error")
	errResponseUnmarshal  = errors.New("error unmarshalling response")
)
