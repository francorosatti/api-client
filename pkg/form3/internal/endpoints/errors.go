package endpoints

import "errors"

var (
	errSerialiseParamValue  = errors.New("error serialising parameter value")
	errBuildUrl             = errors.New("error building url")
	errUnsupportedParamType = errors.New("unsupported parameter type")
	errHttpNewRequest       = errors.New("error creating http request")
	errDoRequest            = errors.New("error doing http request")
)
