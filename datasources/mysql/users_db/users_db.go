package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsersUsername = "MYSQL_USERS_USERNAME"
	mysqlUsersPassword = "MYSQL_USERS_PASSWORD"
	mysqlUsersHost     = "MYSQL_USERS_HOST"
	mysqlUsersSchema   = "MYSQL_USERS_SCHEMA"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	var error error
	Client, error = sql.Open("mysql", datasourceName)

	if error != nil {
		panic(error)
	}

	if error = Client.Ping(); error != nil {
		panic(error)
	}

	log.Println("Database successfully connected")
}
