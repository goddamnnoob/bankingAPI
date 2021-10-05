package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goddamnnoob/notReddit/errs"
	"github.com/goddamnnoob/notReddit/logger"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDb struct {
	client *sqlx.DB
}

func (d UserRepositoryDb) GetAllUsers() ([]User, *errs.AppError) {
	getAll := "select customer_id,name,city,zipcode,date_of_birth, status from customers"
	users := make([]User, 0)
	err := d.client.Select(&users, getAll)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error while selecting Users")
	}
	return users, nil
}

func NewUserRepositoryDb(dbClient *sqlx.DB) UserRepositoryDb {
	return UserRepositoryDb{dbClient}
}

func (d UserRepositoryDb) ById(id string) (*User, *errs.AppError) {
	byId := "select customer_id,name,city,zipcode,date_of_birth, status from customers where customer_id=?"
	var u User
	err := d.client.Get(&u, byId, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customers" + err.Error())
			return nil, errs.NewNotFoundError("User Not Found")
		} else {
			logger.Error("Error while scan user: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected db error")
		}
	}
	return &u, nil
}

func (d UserRepositoryDb) ByStatus(status int) ([]User, *errs.AppError) {
	getUsers := "select customer_id,name,city,zipcode,date_of_birth, status from customers where status=?"
	users := make([]User, 0)
	err := d.client.Select(&users, getUsers, status)
	if err != nil {
		logger.Error("Error while Querying User Table" + err.Error())
		return nil, errs.NewUnexpectedError("Error Querying User Table ")
	}
	return users, nil
}
