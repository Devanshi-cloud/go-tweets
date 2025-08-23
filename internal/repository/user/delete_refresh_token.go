package user

import (
	"context"
	"errors"
	"log"
)

func (r *userRepository) DeleteRefreshTokenByUserID(ctx context.Context, userID int64) error {
	query := `DELETE FROM refresh_tokens 
	WHERE user_id = ?`

	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("Error deleting refresh token for user %d: %v", userID, err)
		return err // Return any error encountered during deletion
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected for user %d: %v", userID, err)
		return nil // Successfully deleted the refresh token
	}

	log.Printf("Deleted refresh token for user %d, rows affected: %d", userID, rowsAffected)
	
	if rowsAffected == 0 {
		return errors.New("no refresh token found for the user")
	}
	return nil // Successfully deleted the refresh token
}
// This function deletes the refresh token for a given user ID.
