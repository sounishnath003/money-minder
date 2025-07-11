package utility

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// GetStringValueFromEnv helps to get the Key from the os Environment
func GetStringValueFromEnv(key, fallback string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	log.Printf("key=%s is not found in os. setting to fallback value", key)
	return fallback
}

// GetNumberValueFromEnv helps to get the Key from the os Environment
func GetNumberValueFromEnv(key string, fallback int) int {
	if value, found := os.LookupEnv(key); found {
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Printf("error parsing environment variable %s: %v", key, err)
			return fallback
		}
		return int(intValue)
	}
	log.Printf("key=%s is not found in os. setting to fallback value", key)
	return fallback
}

// HashPassword hashes a plain password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with a plain password
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// GetJWTSecret returns the JWT secret from env or fallback
func GetJWTSecret() string {
	return GetStringValueFromEnv("JWT_SECRET", "supersecretjwtkey")
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID int, name string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"name":   name,
		"exp":    time.Now().Add(time.Minute * 10).Unix(), // 10 Minute window
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetJWTSecret()))
}

// ValidateJWT validates a JWT token and returns the claims if valid
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(GetJWTSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
