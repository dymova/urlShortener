package models

import (
	"database/sql"
	"fmt"
	"urlShortener/auth"
)

type User struct {
	id       int
	login    string
	password string
}

func SaveUser(login string, password string) error {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}
	_, err = DB.Exec("INSERT INTO users (login, password) VALUES (?, ?)", login, hash)
	if err != nil {
		return err
	}
	return nil
}

func LoginCheck(login string, password string) (string, error) {
	var user User
	row := DB.QueryRow("SELECT password from users where login=?", login)
	if err := row.Scan(&user.login, &user.password); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("LoginCheck %v: no such user", login)
		}
		return "", fmt.Errorf("LoginCheck %v: %v", login, err)
	}

	if auth.VerifyPassword(password, user.password) == false {
		return "", fmt.Errorf("LoginCheck %v: invalid password", login)
	}

	//todo use id instead login
	token, err := auth.IssueJWTToken(user.login)

	if err != nil {
		return "", err
	}

	return token, nil
}
