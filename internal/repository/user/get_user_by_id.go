package user

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (*model.UserModel, error) {
	query := `SELECT id, email, username, created_at, updated_at
	FROM users 
	WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, userID)
	var result model.UserModel
	err := row.Scan(&result.ID, &result.Email, &result.Username, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given ID
		}
		return nil, err // Return any other error encountered
	}
	return &result, nil // Return the found user
}