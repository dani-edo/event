package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "secret" // temporary

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // check is the method is correct
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("Invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims) // check what kind of claims type
	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }

	// email := claims["email"].(string) // is the value type is string?
	// userId := claims["userId"].(int64)
	return nil
}
