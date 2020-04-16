package repo

import (
	"log"
)

func Update(Id, Name, Email, Password, Note string) error {
	log.Println("-------------------------Update----------------------------")
	db, err := connect()
	if err != nil {
		return err
	}

	sql := "UPDATE user SET name = '" + Name + "', email= '" + Email + "', password = '" + Password + "', note = '" + Note + "' WHERE id =" + Id
	log.Println("sqlStr = " + sql)
	_, err = db.Query(sql)

	if err != nil {
		log.Fatal(db)
	}
	return err
}
