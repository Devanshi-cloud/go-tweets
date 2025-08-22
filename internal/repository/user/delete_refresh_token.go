package user

import (
	"context"
	"errors"
)

func (r *userRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID int64) error {
	query := `DELETE FROM refresh_tokens 
	WHERE user_id = ?`

	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return err // Return any error encountered during deletion
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {		
		return nil // Successfully deleted the refresh token
	}

	if rowsAffected == 0 {
		return errors.New("no refresh token found for the user")
	}
	return nil // Successfully deleted the refresh token
}
// This function deletes the refresh token for a given user ID.