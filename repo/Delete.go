package repo

import (
	"log"
)

func Delete(Id string) error {
	log.Println("-------------------------Delete----------------------------")
	db, err := connect()
	if err != nil {
		return err
	}

	sql := "DELETE FROM user WHERE id = " + Id
	log.Println("sqlStr = " + sql)
	_, err = db.Query(sql)

	if err != nil {
		log.Fatal(db)
	}
	return err
}
