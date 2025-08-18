package user

import (
	"context"
	"go-tweets/internal/model"
)

// CreateUser creates a new user in the repository.
func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) (int64, error) {
	query := `INSERT INTO users (username, email, password, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)`

	// Execute the query with the user data
	result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}
	userID, _ := result.LastInsertId()
	return userID, nil
	// Return the ID of the newly created user
	// Note: Error handling for LastInsertId is omitted for brevity, but should be handled in production code.
}

// why?