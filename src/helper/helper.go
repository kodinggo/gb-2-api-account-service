package helper

import (
	"account-service/src/config"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
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

func sanitizeFilename(filename string) string {
	filename = filepath.Base(filename)

	filename = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '.' || r == '-' || r == '_' {
			return r
		}
		return -1
	}, filename)

	return filename
}

func SaveUploadedFile(fileHeader *multipart.FileHeader, uploadPath string) (string, error) {
	if fileHeader == nil {
		return "", errors.New("file header is nil")
	}
	if uploadPath == "" {
		return "", errors.New("upload path is empty")
	}

	filename := sanitizeFilename(fileHeader.Filename)
	if filename == "" {
		return "", errors.New("invalid filename")
	}

	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", err
	}

	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	fullPath := filepath.Join(uploadPath, filename)
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		os.Remove(fullPath)
		return "", err
	}

	return fullPath, nil
}
