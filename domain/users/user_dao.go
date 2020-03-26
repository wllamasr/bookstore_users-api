package users

import (
	"fmt"

	"github.com/wllamasr/bookstore_users-api/datasources/mysql/users_db"

	"github.com/wllamasr/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {

	if error := users_db.Client.Ping(); error != nil {
		panic(error)
	}

	result := usersDB[user.Id]

	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {

	smtm, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer smtm.Close()

	insertResult, error := smtm.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if error != nil {
		return errors.InternalServerError(fmt.Sprintf("error trying to save user: %s", error.Error()))
	}

	userId, error := insertResult.LastInsertId()

	if error != nil {
		return errors.InternalServerError(fmt.Sprintf("error trying to save user: %s", error.Error()))
	}

	user.Id = userId

	usersDB[user.Id] = user
	return nil
}
