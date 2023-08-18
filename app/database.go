package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"golang-fiber/helper"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "go-fiber"
)

func Database() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicError(err) // Menggunakan helper.PanicError untuk menangani error

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(25)
	db.SetConnMaxLifetime(60 * time.Minute) // 1 jam
	db.SetConnMaxIdleTime(10 * time.Minute) // 10 menit

	return db
}
