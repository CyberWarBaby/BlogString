package models

import (
	"time"

	"example.com/blog-api/db"
)

type Blog struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	AuthorID  int64     `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAllBlogs() ([]Blog, error) {
	query := "SELECT * FROM blogs"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var blog Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt)

		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}
