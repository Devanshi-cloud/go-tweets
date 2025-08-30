package post

import "context"

func(r *postRepository) DeleteLikePost(ctx context.Context, postID, userID int64) error {
	query := `DELETE FROM post_likes 
	WHERE post_id = ?
	AND user_id = ?` //simply delete the row??
	_, err := r.db.ExecContext(ctx, query, postID, userID) //sql.Result only tells how many rows were affected
	 //ExecContext is method from database/sql package
	
	return err
}