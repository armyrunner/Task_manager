package models

import "github.com/golang-jwt/jwt/v5"

// =============================================================================
// Task Models
// =============================================================================

// Task represents a task in the system
type Task struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id,omitempty"`
	OriginalID   int    `json:"original_id,omitempty"`
	CategoryID   int    `json:"category_id,omitempty"`
	CategoryName string `json:"category_name,omitempty"` // For display purposes (joined from categories table)
	Description  string `json:"description"`
	DueDate      string `json:"due_date"`
	StartDate    string `json:"start_date"`
	FinishDate   string `json:"finish_date"`
	Status       string `json:"status"`
	Notes        string `json:"notes"`
}

// =============================================================================
// Category Models
// =============================================================================

// Category represents a task category
type Category struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
}

// =============================================================================
// User Models
// =============================================================================

// User represents a user in the system
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // Never send password hash to client
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
}

// =============================================================================
// Auth Request/Response Models
// =============================================================================

// RegisterRequest is the request body for user registration
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest is the request body for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthResponse is returned after successful login/register
type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// TokenRefreshRequest is the request body for refreshing tokens
type TokenRefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// =============================================================================
// Token Models
// =============================================================================

// RefreshToken represents a refresh token in the database
type RefreshToken struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}

// JWTClaims extends jwt.RegisteredClaims with custom user data
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
