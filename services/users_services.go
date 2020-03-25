package services

import (
	"github.com/wllamasr/bookstore_users-api/domain/users"

	"github.com/wllamasr/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}

	if error := result.Get(); error != nil {
		return nil, error
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
