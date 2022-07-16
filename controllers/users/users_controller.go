package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "need implementation",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "need implementation on get",
	})
}

func GetUserById(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "need implementation on get by id",
	})
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
