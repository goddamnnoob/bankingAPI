package domain

import (
	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

type Account struct {
	AccountId   string
	UserId      string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(string) (*Account, *errs.AppError)
	SaveTransactions(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{a.AccountId}
}
