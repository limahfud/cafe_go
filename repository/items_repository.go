package repository

import (
	"cafe_go/model"
	"context"
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

func FindAllAvailableItems() []model.AvailableItem {
	db := GetDBConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, slug, name, description, price FROM items"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	items := []model.AvailableItem{}
	for rows.Next() {
		var id int
		var slug, name, description string
		var price float32

		err := rows.Scan(&id, &slug, &name, &description, &price)
		if err != nil {
			panic(err)
		}

		item := model.AvailableItem{
			Slug:        slug,
			Name:        name,
			Description: description,
			Price:       price,
		}

		items = append(items, item)
	}

	return items
}
