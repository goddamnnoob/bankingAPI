package dto

import (
	"strings"

	"github.com/goddamnnoob/notReddit/errs"
)

type TransactionRequest struct {
	UserId          string  `json:"user_id"`
	TransactionType string  `json:"transaction_type"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
}

func (t TransactionRequest) Validate() *errs.AppError {
	if strings.ToLower(t.TransactionType) != "withdraw" && strings.ToLower(t.TransactionType) != "deposit" {
		return errs.NewValidationError("Transaction type is invalid only withdraw or deposit")
	}
	if t.Amount <= 0 {
		return errs.NewValidationError("Withdraw or Deposit amount cannot be negative or zero")
	}
	return nil
}
