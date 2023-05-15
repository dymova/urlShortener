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
	full      string
	shortened string
}

func ShortenUrl(url string) (string, error) {
	//todo handle colisions
	shortened := os.Getenv("BASE_URL") + "/redirect/" + GenerateRandomString()
	_, err := DB.Exec("INSERT INTO urls (long, short) VALUES (?, ?)", url, shortened)
	if err == nil {
		return "", err
	}

	return shortened, nil
}

const length = 10

func GenerateRandomString() string {
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
