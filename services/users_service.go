package services

import (
	user_domain "vandi_users-api/domain/users_domain"
	"vandi_users-api/utils/errors"
)

func GetUser(user *user_domain.User) *errors.RestErr {

	err := user.GetUser()
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(userId int64) *errors.RestErr {

	result := &user_domain.User{Id: userId}
	err := result.GetUserById(userId)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(user user_domain.User) (*user_domain.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	err := user.SaveUser()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
