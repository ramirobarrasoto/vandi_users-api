package app

import (
	"vandi_users-api/controllers/ping"
	"vandi_users-api/controllers/users"
)

func mapsUrl() {
	//for ping
	router.GET("/ping", ping.Ping)
	//for users
	router.GET("/users", users.GetUser)
	router.GET("/users/:id", users.GetUserById)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:id", users.PutUser)
	router.DELETE("/users/:id", users.DeleteUser)
}
