package helper

import (
	"account-service/src/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashRequestPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID int64) (strToken string, err error) {
	expiredAt := time.Now().UTC().Add(config.JWTExp())
	strToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     expiredAt.Unix(),
		"user_id": userID,
	}).SignedString([]byte(config.JWTSigningKey()))
	return
}
