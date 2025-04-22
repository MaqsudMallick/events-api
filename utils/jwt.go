package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(username string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func ValidateToken(token string) (int64, error) {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return 0, errors.New("invalid token")
	}

	if _, ok := claims["exp"]; !ok {
		return 0, errors.New("token is missing an expiration time")
	}

	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		return 0, errors.New("token is expired")
	}
	if _, ok := claims["userId"]; !ok {
		return 0, errors.New("token is missing a userId")
	}
	if _, ok := claims["username"]; !ok {
		return 0, errors.New("token is missing a username")
	}
	return int64(claims["userId"].(float64)), nil
}
