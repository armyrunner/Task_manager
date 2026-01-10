package db

import (
	"database/sql"
	"time"

	"github.com/armyrunner/task_manager/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	stmt, err := DB.Prepare(`
		INSERT INTO users (username, email, password_hash, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().Format(time.RFC3339)
	result, err := stmt.Exec(user.Username, user.Email, user.PasswordHash, now, now)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	user.CreatedAt = now
	user.UpdatedAt = now

	return nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(email string) (*models.User, error) {
	stmt, err := DB.Prepare(`
		SELECT id, username, email, password_hash, created_at, updated_at 
		FROM users WHERE email = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id int) (*models.User, error) {
	stmt, err := DB.Prepare(`
		SELECT id, username, email, password_hash, created_at, updated_at 
		FROM users WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (*models.User, error) {
	stmt, err := DB.Prepare(`
		SELECT id, username, email, password_hash, created_at, updated_at 
		FROM users WHERE username = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

// SaveRefreshToken stores a refresh token in the database
func SaveRefreshToken(userID int, token string, expiresAt time.Time) error {
	stmt, err := DB.Prepare(`
		INSERT INTO refresh_tokens (user_id, token, expires_at, created_at) 
		VALUES (?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, token, expiresAt.Format(time.RFC3339), time.Now().Format(time.RFC3339))
	return err
}

// GetRefreshToken retrieves a refresh token from the database
func GetRefreshToken(token string) (*models.RefreshToken, error) {
	stmt, err := DB.Prepare(`
		SELECT id, user_id, token, expires_at, created_at 
		FROM refresh_tokens WHERE token = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rt models.RefreshToken
	err = stmt.QueryRow(token).Scan(
		&rt.ID,
		&rt.UserID,
		&rt.Token,
		&rt.ExpiresAt,
		&rt.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Token not found
		}
		return nil, err
	}

	return &rt, nil
}

// DeleteRefreshToken removes a refresh token from the database
func DeleteRefreshToken(token string) error {
	stmt, err := DB.Prepare("DELETE FROM refresh_tokens WHERE token = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	return err
}

// DeleteUserRefreshTokens removes all refresh tokens for a user (logout all devices)
func DeleteUserRefreshTokens(userID int) error {
	stmt, err := DB.Prepare("DELETE FROM refresh_tokens WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	return err
}
