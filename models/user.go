package models

import (
	"database/sql"
	"fmt"
	"urlShortener/auth"
)

type User struct {
	Login    string
	Password string
}

func (u *User) SaveUser() error {
	hash, err := auth.HashPassword(u.Password)
	if err != nil {
		return err
	}
	_, err = DB.Exec("INSERT INTO users (login, password) VALUES (?, ?, ?)", u.Login, hash)
	if err != nil {
		return err
	}
	return nil
}

func LoginCheck(login string, password string) (string, error) {
	var user User
	row := DB.QueryRow("SELECT password from users where login=?", login)
	if err := row.Scan(&user.Login, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("LoginCheck %v: no such user", login)
		}
		return "", fmt.Errorf("LoginCheck %v: %v", login, err)
	}

	if auth.VerifyPassword(password, user.Password) == false {
		return "", fmt.Errorf("LoginCheck %v: invalid password", login)
	}

	//todo use id instead login
	token, err := auth.IssueJWTToken(user.Login)

	if err != nil {
		return "", err
	}

	return token, nil
}
