package store

import (
	"database/sql"
	"errors"
	"log"
	"phonebook/model"
)

func Save(db *sql.DB, user model.User) {
	_, err := db.Exec("INSERT into registration VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		log.Fatal(err)
	}
}

func Retrieve(db *sql.DB, user model.User) (model.User, error) {
	var us model.User

	q := "SELECT * FROM registration WHERE email = '" + user.Email + "';"
	err := db.QueryRow(q).Scan(&us.Email, &us.Password)
	if err != nil{
		log.Fatal(err)
	}

	if us.Email == "" {
		err = errors.New("ErrNotFound")
		return us, err
	}

	if us.Password != user.Password {
		err = errors.New("ErrWrongPass")
		return us, err
	}

	return us, err
}
