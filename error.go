package bri

import (
	"errors"
)

// ErrPendingTransaction defines error if BRI response with http status 200 but html error body.
// Transaction should be pending and need to be inquired.
var ErrPendingTransaction = errors.New("Transaction is pending")
