package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/armyrunner/task_manager/models"
	"github.com/golang-jwt/jwt/v5"
)

// Secret key for signing JWTs - In production, use environment variable!
var JWTSecret = []byte("your-secret-key-change-in-production")

// Token expiration times
const (
	AccessTokenExpiry  = 15 * time.Minute   // Access token expires in 15 minutes
	RefreshTokenExpiry = 7 * 24 * time.Hour // Refresh token expires in 7 days
)

// GenerateAccessToken creates a new JWT access token for a user
func GenerateAccessToken(user *models.User) (string, error) {
	claims := models.JWTClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "task-manager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// GenerateRefreshToken creates a random refresh token
func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// ValidateAccessToken validates a JWT access token and returns the claims
func ValidateAccessToken(tokenString string) (*models.JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetRefreshTokenExpiry returns the expiration time for a refresh token
func GetRefreshTokenExpiry() time.Time {
	return time.Now().Add(RefreshTokenExpiry)
}
