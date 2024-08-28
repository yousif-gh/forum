package models

type Comment struct {
	ID        int    `json:"id"`
	PostID    int    `json:"post_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Username  string `json:"username"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislikes"`
}
