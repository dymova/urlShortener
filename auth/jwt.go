package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func IssueJWTToken(user int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secret := os.Getenv("API_JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractJWTToken(tokenString string) (*jwt.Token, error) {
	// validate token and exp data
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid != true {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}

func ExtractJWTTokenUser(tokenString string) (int, error) {
	token, err := ExtractJWTToken(tokenString)
	if err != nil {
		return -1, err
	}
	claims, isValid := token.Claims.(jwt.MapClaims)
	if isValid != true {
		return -1, fmt.Errorf("invalid token")
	} else {
		return int(claims["user_id"].(float64)), nil
	}
}
