package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() *sql.DB {
	driverName := "mysql"
	userName := "freedb_hojlund"
	password := "W2MDvxGzSknCh#E"
	host := "sql.freedb.tech"
	port := 3306
	dbName := "freedb_superdb"
	urlParam := "?parseTime=true"

	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s", userName, password, host, port, dbName, urlParam)
	db, err := sql.Open(driverName, url)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
