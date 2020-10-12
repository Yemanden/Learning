package models

import "errors"

type (
	AdapterError    error
  SingletonError  error
)

var (
	DocumentMissingDataError AdapterError     = errors.New("the document is missing data")  	
  SellerMissingProductError SingletonError  = errors.New("product is missing")
)
