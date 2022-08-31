package users_domain

import (
	"fmt"
	"strings"
	"vandi_users-api/datasources/mysql/users_db"
	"vandi_users-api/utils/date_utils"
	"vandi_users-api/utils/errors"
)

const (
	queryInsertUser   = "INSERT INTO users(first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUsers     = "SELECT * FROM users;"
	queryGetUsersById = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

//here we have the conection with the DB

var usersDB = make(map[int64]*User)

func (user *User) GetUser() *errors.RestErr {

	statement, err := users_db.Client.Prepare(queryGetUsers)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	result, err := statement.Query()

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to get users: %s", err.Error()),
		)
	}
	fmt.Sprint(result)

	/*
		if usersDB[user.Id] == nil {
			return nil, nil
		}
	*/

	return nil
}

func (user *User) GetUserById(id int64) *errors.RestErr {

	statement, err := users_db.Client.Prepare(queryGetUsersById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	result := statement.QueryRow(user.Id)

	if err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to get user %d: %s", &user.Id, err.Error()),
		)
	}

	return nil
	/*
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
	*/
}

func (user *User) PutUser(id int64) (*User, *errors.RestErr) {
	return nil, nil
}

func (user *User) SaveUser() *errors.RestErr {

	statement, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer statement.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), "users.email_UNIQUE") {
			return errors.NewBadRequestError(
				fmt.Sprintf("The email %s is already exist", user.Email),
			)
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to save the user: %s", err.Error()),
		)
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error when trying to save the user: %s", err.Error()),
		)
	}

	user.Id = userId
	/*
		current := usersDB[user.Id]

		if current != nil {
			if current.Email == user.Email {
				return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
			}
			return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
		}

		user.DateCreated = date_utils.GetNowString()
		usersDB[user.Id] = user
	*/
	return nil
}

func (user *User) DeleteUser(id int64) *errors.RestErr {
	return nil
}
