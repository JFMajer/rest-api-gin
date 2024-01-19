package models

import (
	"github.com/JFMajer/rest-api-gin/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() (int, error) {
	passwordHash, err := hashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	query :=
		`INSERT INTO users (email, password)
			VALUES (?, ?)`

	statement, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	result, err := statement.Exec(u.Email, passwordHash)
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
