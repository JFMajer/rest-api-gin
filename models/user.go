package models

import (
	"github.com/JFMajer/rest-api-gin/db"
)

type User struct {
	ID       int64
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u *User) Save() (int, error) {
	query :=
		`INSERT INTO users (email, password)
			VALUES (?, ?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(u.email, u.password)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	u.ID = lastId

	return int(lastId), nil

}
