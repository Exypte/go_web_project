package swagger

import (
	"fmt"
	"context"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

var (
	ctx context.Context
	db *sql.DB
)

func Connect(user, password, dbname string) *sql.DB{
	connectionString := fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", user, password, dbname)
	
	var err error

	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(connectionString)

	return db
}

func GetDB() *sql.DB{
	return db
}