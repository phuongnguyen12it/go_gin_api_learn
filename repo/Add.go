package repo

import (
	"log"
)

func Add(Name, Email, Password, Note string) error {
	log.Println("-------------------------Add----------------------------")
	db, err := connect()
	if err != nil {
		return err
	}

	sql := "INSERT INTO user (name, email, password, note) VALUES('" + Name + "', '" + Email + "', '" + Password + "', '" + Note + "')"
	_, err = db.Query(sql)

	if err != nil {
		log.Fatal(db)
	}
	return err
}
