package mysql

import (
	"account-service/account-service/src/config"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnStr() string {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.MySQLDBUser(),
		config.MySQLDBPass(),
		config.MySQLDBHost(),
		config.MySQLDBPort(),
		config.MySQLDBName())

	fmt.Println("Connection string:", connStr)
	return connStr
}

func InitDBConn() *sql.DB {
	dbConn, err := sql.Open("mysql", NewConnStr())
	if err != nil {
		panic(err)
	}

	dbConn.SetConnMaxLifetime(time.Minute * 3)
	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	if err := dbConn.Ping(); err != nil {
		panic(err)
	}

	return dbConn
}
