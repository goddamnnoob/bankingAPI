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

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
