package security

import (
	"devbook/src/config"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetToken generates a jwt token with user permissions
func GetToken(userId uint64, userEmail string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userId
	claims["userEmail"] = userEmail

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

// ValidateToken validates a request JWT token
func ValidateToken(r *http.Request) error {
	stringToken := extractToken(r)
	token, error := jwt.Parse(stringToken, getVerificationKey)
	if error != nil {
		return error
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Invalid Token")
}

func ExtractUserId(r *http.Request) (uint64, error) {
	stringToken := extractToken(r)
	token, error := jwt.Parse(stringToken, getVerificationKey)
	if error != nil {
		return 0, error
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		userId, error := strconv.ParseUint(
			fmt.Sprintf("%.0f",claims["userId"]), 10, 64)

		if error != nil {
			return 0, error
		}
		return userId, nil
	}
	return 0, errors.New("Invalid Token")
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Signing Method Unespected. %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}