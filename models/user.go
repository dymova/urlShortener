package models

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"urlShortener/auth"
)

type User struct {
	Id       int
	Login    string
	Password string
}

func SaveUser(login string, password string) error {
	hash, err := hashPassword(password)
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
	row := DB.QueryRow("SELECT id, login, password from users where login = ?", login)
	if err := row.Scan(&user.Id, &user.Login, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("LoginCheck %v: no such user", login)
		}
		return "", fmt.Errorf("LoginCheck %v: %v", login, err)
	}

	if verifyPassword(password, user.Password) == false {
		return "", fmt.Errorf("LoginCheck %v: invalid password", login)
	}

	token, err := auth.IssueJWTToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(id int) (User, error) {
	var user User
	row := DB.QueryRow("SELECT id, login, password from users where id = ?", id)
	if err := row.Scan(&user.Id, &user.Login, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("GetUser %v: no such user", id)
		}
		return user, fmt.Errorf("GetUser %v: %v", id, err)
	}
	return user, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
