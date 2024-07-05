package auth_utils

import (
	"BeatBoxBox/pkg/logger"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const DEFAULT_TOKEN_EXPIRATION = 48 * time.Hour

// HashString hashes a password using bcrypt
func HashString(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	for attempts := 0; err != nil && attempts < 20; attempts++ {
		hashed_password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	}
	return string(hashed_password), err
}

// GenerateToken generates a random token of 128 bits
func GenerateRandomTokenWithHash() (string, string, error) {
	bytes := make([]byte, 36)
	if _, err := rand.Read(bytes); err != nil {
		return "", "", err
	}
	raw_token := hex.EncodeToString(bytes)
	hashed_token, err := HashString(raw_token)
	if err != nil {
		return "", "", err
	}
	return raw_token, hashed_token, nil
}

// CompareHash compares a hashed password with a plaintext password
func CompareHash(hashed_string string, attempt_string string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed_string), []byte(attempt_string)) == nil
}

func CreateAuthJWT(user_id int, auth_token string) (string, error) {
	JWT_SECRET := os.Getenv("SECRET_JWT_KEY")
	if JWT_SECRET == "" {
		logger.Critical("Missing environment variable SECRET_JWT_KEY")
		return "", errors.New("missing environment variable SECRET_JWT_KEY")
	}
	claims := jwt.MapClaims{
		"user_id":    user_id,
		"auth_token": auth_token,
		"expiration": time.Now().Add(DEFAULT_TOKEN_EXPIRATION).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET))
}

func ParseAuthJWT(token_string string) (int, string, error) {
	JWT_SECRET := os.Getenv("SECRET_JWT_KEY")
	if JWT_SECRET == "" {
		logger.Critical("Missing environment variable SECRET_JWT_KEY")
		return -1, "", errors.New("missing environment variable SECRET_JWT_KEY")
	}
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(JWT_SECRET), nil
	})

	user_id := -1
	auth_token := ""

	if err != nil {
		return user_id, auth_token, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id = int(claims["user_id"].(float64))
		auth_token = claims["auth_token"].(string)
		expiration_time := int64(claims["expiration"].(float64))
		if time.Now().Unix() > expiration_time {
			return user_id, auth_token, errors.New("token expired")
		}
		return user_id, auth_token, nil
	}

	return user_id, auth_token, errors.New("invalid token")
}
