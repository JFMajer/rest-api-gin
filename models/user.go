package models

import (
	"database/sql"
	"errors"

	"github.com/JFMajer/rest-api-gin/db"
	"github.com/JFMajer/rest-api-gin/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() (int, error) {
	passwordHash, err := utils.HashPassword(u.Password)
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

func GetUsers() ([]User, error) {
	query :=
		`SELECT id, email FROM users`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *User) ValidateCredentials() error {
	query :=
		`SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			// User does not exist
			return errors.New("invalid credentials")
		}
		// Other error
		return errors.New("invalid credentials")
	}

	passOK := utils.VerifyPassword(retrievedPassword, u.Password)
	if !passOK {
		return errors.New("invalid credentials")
	}

	return nil
}
