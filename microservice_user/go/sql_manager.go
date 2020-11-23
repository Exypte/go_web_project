package swagger

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect(){
	db, err := sql.Open("postgres","postgres://root:pwd@localhost:5432/mydb?sslmode=verify-full")

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection ok")
	defer db.Close()
}