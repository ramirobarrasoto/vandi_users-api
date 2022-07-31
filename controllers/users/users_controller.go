package users

import (
	"net/http"
	"vandi_users-api/domain/users_domain"
	"vandi_users-api/services"
	"vandi_users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

var user users_domain.User
var id int64

func GetUser(c *gin.Context) {

	getErr := services.GetUser(&user)

	if getErr != nil {
		//TODO: handler the error please
		return
	}

	c.JSON(http.StatusOK, nil)
}

func GetUserById(c *gin.Context) {
	getErr := services.GetUserById(id)

	if getErr != nil {
		//TODO: handler the error
		return
	}

	c.JSON(http.StatusCreated, nil)
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
