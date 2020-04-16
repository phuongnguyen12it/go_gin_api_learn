package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() (db *sql.DB, err error) {
	log.Println("-------------------------connect db----------------------------")
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang_demo"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)

	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
