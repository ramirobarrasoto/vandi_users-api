package users

import (
	"net/http"
	"strconv"
	"vandi_users-api/domain/users_domain"
	"vandi_users-api/services"
	"vandi_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

var user users_domain.User
var id int64

func GetUser(c *gin.Context) {

	getData, getErr := services.GetUser(&user)

	if getErr != nil {
		//TODO: handler the error please
		return
	}

	c.JSON(http.StatusOK, getData)
}

func GetUserById(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("invalid format, user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUserById(userId)
	if getErr != nil {
		err := errors.NewNotFoundError("user not found")
		c.JSON(getErr.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
func CreateUser(c *gin.Context) {

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, createErr := services.CreateUser(user)

	if createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func PutUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "need implementation on put",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "need implementation on delete",
	})
}
