package models

import (
	"database/sql"
	"time"

	"example.com/blog-api/db"
	"example.com/blog-api/utils"
)

type Blog struct {
	ID        int64
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthorID  int64
	Slug      string
}

func (b *Blog) Save() error {

	b.AuthorID = 1
	b.Slug = utils.GenSlug(b.Title)
	query := `INSERT INTO blogs (title, content, author_id, slug) VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Content, b.AuthorID, b.Slug)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	b.ID = id

	row := db.DB.QueryRow("SELECT created_at, updated_at FROM blogs WHERE id = ?", b.ID)
	var createdAt, updatedAt string
	if err := row.Scan(&createdAt, &updatedAt); err != nil {
		return err
	}

	b.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	b.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAt)

	return nil
}

func GetAllBlogs() ([]Blog, error) {
	query := "SELECT id, title, slug, content, author_id, created_at, updated_at FROM blogs"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var blog Blog
		var authorID sql.NullInt64 // <-- handle possible NULL
		var createdAtStr, updatedAtStr string

		err := rows.Scan(&blog.ID, &blog.Title, &blog.Slug, &blog.Content, &authorID, &createdAtStr, &updatedAtStr)
		if err != nil {
			return nil, err
		}

		if authorID.Valid {
			blog.AuthorID = authorID.Int64
		} else {
			blog.AuthorID = 0 // default for NULL
		}

		blog.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
		blog.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

		blogs = append(blogs, blog)
	}

	return blogs, nil
}

// func GetBlogById(id int64) (*Blog, error) {
// 	query := "SELECT id, title, content, author_id, created_at, updated_at FROM blogs WHERE id = ?"

// 	row := db.DB.QueryRow(query, id)
// 	var blog Blog

// 	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
// 		return nil, err
// 	}
// 	return &blog, nil

// }

func GetBlogById(id int64) (*Blog, error) {
	query := "SELECT id, title, content, author_id, created_at, updated_at FROM blogs WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var blog Blog
	var authorID sql.NullInt64
	var createdAtStr, updatedAtStr string

	// scan safely
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &authorID, &createdAtStr, &updatedAtStr); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no blog found
		}
		return nil, err
	}

	// handle possible NULL author_id
	if authorID.Valid {
		blog.AuthorID = authorID.Int64
	} else {
		blog.AuthorID = 0 // default
	}

	// parse timestamps from sqlite strings
	blog.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
	blog.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

	return &blog, nil
}
