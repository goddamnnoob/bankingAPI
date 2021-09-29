package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goddamnnoob/notReddit/errs"
	"github.com/goddamnnoob/notReddit/logger"
)

type UserRepositoryDb struct {
	client *sql.DB
}

func (d UserRepositoryDb) GetAllUsers() ([]User, *errs.AppError) {
	getAll := "select customer_id,name,city,zipcode,date_of_birth, status from customers"
	rows, err := d.client.Query(getAll)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error Querying Customer Table ")
	}
	users := make([]User, 0)
	for rows.Next() {

		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error while scanningUsers")
		}
		users = append(users, u)
	}
	return users, nil
}

func NewUserRepositoryDb() UserRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxIdleConns(10)
	client.SetMaxOpenConns(10)
	return UserRepositoryDb{client}
}

func (d UserRepositoryDb) ById(id string) (*User, *errs.AppError) {
	byId := "select customer_id,name,city,zipcode,date_of_birth, status from customers where customer_id=?"
	rows := d.client.QueryRow(byId, id)
	var u User
	err := rows.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)
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
	rows, err := d.client.Query(getUsers, status)
	if err != nil {
		logger.Error("Error while Querying User Table" + err.Error())
		return nil, errs.NewUnexpectedError("Error Querying User Table ")
	}
	users := make([]User, 0)
	for rows.Next() {

		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)
		if err != nil {
			return nil, errs.NewUnexpectedError("Error while scanningUsers")
		}
		users = append(users, u)
	}
	return users, nil
}
