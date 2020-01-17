package models

import "errors"

type (
	AdapterError error
)

var (
	DocumentMissingDataError AdapterError = errors.New("the document is missing data")
)
