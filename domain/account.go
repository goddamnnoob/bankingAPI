package domain

import (
	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

type Account struct {
	AccountId   string  `json:"account_id"`
	UserId      string  `json: "user_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
	SaveTransactions(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}
