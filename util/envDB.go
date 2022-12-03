package util

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

func EnvDB() *sqlx.DB {
	env := EnvCheck()
	dbUser := env.DB.User
	dbPasswd := env.DB.Password
	dbAddr := env.DB.Address
	dbPort := env.DB.Port
	dbName := env.DB.Name

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
