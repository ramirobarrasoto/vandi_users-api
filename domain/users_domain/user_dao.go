package users_domain

import (
	"fmt"
	"vandi_users-api/utils/date_utils"
	"vandi_users-api/utils/errors"
)

//here we have the conection with the DB

var usersDB = make(map[int64]*User)

func (user *User) GetUser() (map[int64]*User, *errors.RestErr) {

	if usersDB[user.Id] == nil {
		return nil, nil
	}

	return usersDB, nil
}

func (user *User) GetUserById(id int64) (*User, *errors.RestErr) {
	result := usersDB[user.Id]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return result, nil
}

func (user *User) PutUser(id int64) (*User, *errors.RestErr) {
	return nil, nil
}

func (user *User) SaveUser() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}

func (user *User) DeleteUser(id int64) *errors.RestErr {
	return nil
}
