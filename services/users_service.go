package services

import (
	user_domain "vandi_users-api/domain/users_domain"
	"vandi_users-api/utils/errors"
)

func GetUser(user *user_domain.User) (map[int64]*user_domain.User, *errors.RestErr) {

	result, err := user.GetUser()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUserById(userId int64) (*user_domain.User, *errors.RestErr) {

	result := &user_domain.User{Id: userId}
	user, err := result.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
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
