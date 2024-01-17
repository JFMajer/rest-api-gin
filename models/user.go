package models

import (
	"github.com/JFMajer/rest-api-gin/db"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
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
	result, err := statement.Exec(u.Email, u.Password)
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
