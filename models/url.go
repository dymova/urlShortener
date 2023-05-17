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
	Full      string
	ShortCode string
	owner     string
}

func ShortenUrl(url string, userId int) (string, error) {
	//todo handle colisions
	shortCode := generateRandomString()
	_, err := DB.Exec("INSERT INTO urls (full, shortCode, owner) VALUES (?, ?, ?)", url, shortCode, userId)
	if err != nil {
		return "", err
	}

	return os.Getenv("BASE_URL") + "/redirect/" + shortCode, nil
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
	row := DB.QueryRow("SELECT full FROM urls WHERE shortCode=?", shortCode)
	if err := row.Scan(&full); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("GetFullUrl %v: no such url", shortCode)
		}
		return "", fmt.Errorf("GetFullUrl %v: %v", shortCode, err)
	}
	return full, nil
}

func GetUsersUrls(user User) ([]Url, error) {
	// An urls slice to hold data from returned rows.
	var urls []Url

	rows, err := DB.Query("SELECT id, full, shortCode, owner  FROM urls WHERE owner = ?", user.Id)
	if err != nil {
		return nil, fmt.Errorf("GetUsersUrls %q: %v", user, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var url Url
		if err := rows.Scan(&url.id, &url.Full, &url.ShortCode, &url.owner); err != nil {
			return nil, fmt.Errorf("GetUsersUrls %q: %v", user, err)
		}
		urls = append(urls, url)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetUsersUrls %q: %v", user, err)
	}
	return urls, nil
}
