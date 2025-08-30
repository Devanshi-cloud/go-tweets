package comment

import (
	"context"
	"go-tweets/internal/model"
)

func (r *commentRepository) StoreComment(ctx context.Context, model *model.CommentModel) error {
	query := `INSERT INTO comments (id, user_id, tweet_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, model.ID, model.UserID, model.Content, model.CreatedAt, model.UpdatedAt)

	return err
}