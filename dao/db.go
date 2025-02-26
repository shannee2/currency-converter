package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "shannee_ahirwar_ftc"
	password = "postgres"
	dbname   = "currency"
)

var db *sql.DB

func InitDB() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully!")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
