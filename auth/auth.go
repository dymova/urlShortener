package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func IssueJWTToken(user string) (string, error) {
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

func ExtractJWTTokenUser(tokenString string) (string, error) {
	token, err := ExtractJWTToken(tokenString)
	if err != nil {
		return "", err
	}
	claims, isValid := token.Claims.(jwt.MapClaims)
	if isValid != true {
		return "", fmt.Errorf("invalid token")
	} else {
		user := claims["user_id"].(string)
		//todo check that such user exist
		return user, nil
	}
}
