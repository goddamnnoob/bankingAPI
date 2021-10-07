package dto

import (
	"strings"

	"github.com/goddamnnoob/notReddit/errs"
)

type NewAccountRequest struct {
	UserId      string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need atleast 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account Type is invalid")
	}
	return nil
}
