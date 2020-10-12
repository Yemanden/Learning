package models

import "errors"

type (
	SingletonError error
)

var (
	SellerMissingProductError SingletonError = errors.New("product is missing")
)
