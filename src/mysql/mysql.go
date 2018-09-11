package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dbInstance *sql.DB

func SharedDb() *sql.DB {

	if dbInstance != nil {
		return dbInstance
	}

	dbInstance, err := sql.Open("mysql", "root:12345678@/mydatabase")
	if err != nil {
		panic(err)
		return nil
	}
	return dbInstance

}
