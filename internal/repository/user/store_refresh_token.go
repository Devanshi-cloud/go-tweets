package user
import (
	"context"
	
	"go-tweets/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, model *model.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, created_at, expired_at, updated_at) 
	VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.CreatedAt, model.ExpiredAt, model.UpdatedAt)
	
	return err // Successfully stored the refresh token
}