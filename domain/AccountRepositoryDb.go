package domain

import (
	"strconv"

	"github.com/goddamnnoob/notReddit/errs"
	"github.com/goddamnnoob/notReddit/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	insert := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) values (?,?,?,?,?)"
	res, err := d.client.Exec(insert, a.UserId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected db error")
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inserted id for account inserted: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected db error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}

func (d AccountRepositoryDb) SaveTransactions(t Transaction) (*Transaction, *errs.AppError) {
	con, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank acount" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected db error")
	}

	result, _ := con.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.TransactionType == "withdraw" {
		_, err = con.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = con.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		con.Rollback()
		logger.Error("Error while saving transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected db error")
	}

	err = con.Commit()

	if err != nil {
		con.Rollback()
		logger.Error("Error while commiting transaction " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected db error " + err.Error())
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while gitting last inserted id from db " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected db error")
	}

	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
