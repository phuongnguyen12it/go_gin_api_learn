package repo

import (
	"log"
	"user-api/model"
)

func Read(id string) (user model.User, err error) {
	log.Println("-------------------------Read----------------------------")
	db, err := connect()
	if err != nil {
		return user, err
	}

	sql := "SELECT * FROM user WHERE id = " + id
	log.Println("sqltrng = " + sql)
	log.Println(db.Query(sql))
	rows, err := db.Query(sql)
	log.Println(rows)
	if err != nil {
		log.Println(err)
	}
	if rows != nil {
		for rows.Next() {
			user := model.User{}
			err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Note)
			return user, err
		}
	}

	return
}
