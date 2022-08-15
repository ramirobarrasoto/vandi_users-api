package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
	"mysql_users_username":"root",
	"mysql_users_password":"elMejor007$",
	"mysql_users_host":"127.0.0.1:3306",
	"mysql_users_schema":"user_db",
*/
/*
const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)
*/
var Client *sql.DB

func init() {
	datasourcesName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		//username, password, host, schema,
		"root",
		"elMejor007$",
		"127.0.0.1:3306",
		"user_db",
	)
	var err error
	Client, err = sql.Open("mysql", datasourcesName)

	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("data base successfully cofigured")
}
