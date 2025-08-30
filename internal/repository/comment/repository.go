package comment

import (
	"context"
	"database/sql"
	"go-tweets/internal/model"
)

type CommentRepository interface {
	StoreComment(ctx context.Context, model *model.CommentModel) error
}

type commentRepository struct {
	db *sql.DB
}

// constructor
func NewRepository(db *sql.DB) CommentRepository {
	return &commentRepository{
		db: db,
	}
}
