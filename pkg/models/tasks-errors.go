package models

import "errors"

type InputErrors error

var (
	InvalidNumbersError InputErrors = errors.New("invalid numbers")
)
