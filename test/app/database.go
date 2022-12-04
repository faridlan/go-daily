package app

import (
	"database/sql"
	"time"

	"github.com/faridlan/daily/test/helper"
)

func NewConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/daily")
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)

	return db

}
