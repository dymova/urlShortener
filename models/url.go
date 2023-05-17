package models

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Url struct {
	id        string
	full      string
	shortened string
	owner     string
}

func ShortenUrl(url string, user string) (string, error) {
	//todo handle colisions
	shortened := os.Getenv("BASE_URL") + "/redirect/" + generateRandomString()
	_, err := DB.Exec("INSERT INTO urls (full, shortened, owner) VALUES (?, ?, ?)", url, shortened, user)
	if err == nil {
		return "", err
	}

	return shortened, nil
}

const length = 10

func generateRandomString() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var builder strings.Builder
	maxIndex := big.NewInt(int64(len(chars)))

	for i := 0; i < length; i++ {
		randomIndex, _ := rand.Int(rand.Reader, maxIndex)
		builder.WriteByte(chars[randomIndex.Int64()])
	}

	return builder.String()
}

func GetFullUrl(shortCode string) (string, error) {
	var full string
	row := DB.QueryRow("SELECT full FROM urls WHERE shortened=?", shortCode)
	if err := row.Scan(&full); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("GetFullUrl %v: no such url", shortCode)
		}
		return "", fmt.Errorf("GetFullUrl %v: %v", shortCode, err)
	}
	return full, nil
}

func GetUsersUrls(user string) ([]Url, error) {
	// An albums slice to hold data from returned rows.
	var albums []Url

	rows, err := DB.Query("SELECT * urls urls WHERE owner = ?", user)
	if err != nil {
		return nil, fmt.Errorf("GetUsersUrls %q: %v", user, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var url Url
		if err := rows.Scan(&url.full, &url.shortened, &url.owner); err != nil {
			return nil, fmt.Errorf("GetUsersUrls %q: %v", user, err)
		}
		albums = append(albums, url)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetUsersUrls %q: %v", user, err)
	}
	return albums, nil
}
