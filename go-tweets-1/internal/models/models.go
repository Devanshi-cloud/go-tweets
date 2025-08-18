package models

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type Tweet struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
	CreatedAt string `json:"created_at"`
}