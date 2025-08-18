package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error) {
	query := `SELECT id, username, email, password, created_at, updated_at 
	FROM users 
	WHERE email = ? 
	OR username = ?`

	row := r.db.QueryRowContext(ctx, query, email, username)
	var result model.UserModel
	if err := row.Scan(&result.ID, &result.Username, &result.Email, &result.Password, &result.CreatedAt, &result.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err // Return any other error encountered
	}
	return &result, nil // Placeholder return, implement actual logic to scan row into UserModel
	// You would typically scan the row into a UserModel instance and return it
}