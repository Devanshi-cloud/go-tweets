package user
import (
	"context"
	"log"
	
	"go-tweets/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, created_at, expired_at, updated_at) 
	VALUES (?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.CreatedAt, model.ExpiredAt, model.UpdatedAt)
	if err != nil {
		log.Printf("Error storing refresh token for user %d: %v", model.UserID, err)
		return err
	}
	
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Successfully stored refresh token for user %d, rows affected: %d", model.UserID, rowsAffected)
	
	return nil // Successfully stored the refresh token
}
