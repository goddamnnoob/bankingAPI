package service

import (
	"time"

	"github.com/goddamnnoob/notReddit/domain"
	"github.com/goddamnnoob/notReddit/dto"
	"github.com/goddamnnoob/notReddit/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountId:   "",
		UserId:      req.UserId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}

	res := newAccount.ToNewAccountResponseDto()

	return &res, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.TransactionType == "withdraw" {
		_, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
	}
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	transaction, appError := s.repo.SaveTransactions(t)
	if err != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
