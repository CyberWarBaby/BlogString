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
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt

	query := `
		INSERT INTO blogs (title, content, author_id, slug, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		b.Title,
		b.Content,
		b.AuthorID,
		b.Slug,
		b.CreatedAt.Format("2006-01-02 15:04:05"),
		b.UpdatedAt.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	b.ID = id

	// now you can scan, but it's optional because you already have the timestamps in struct
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
		// var createdAtStr, updatedAtStr time.Time

		err := rows.Scan(&blog.ID, &blog.Title, &blog.Slug, &blog.Content, &authorID, &blog.CreatedAt, &blog.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if authorID.Valid {
			blog.AuthorID = authorID.Int64
		} else {
			blog.AuthorID = 0 // default for NULL
		}

		// blog.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
		// blog.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

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

	// scan safely
	if err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &authorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
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

	return &blog, nil
}

func (blog Blog) Update() error {
	query := `
	UPDATE blogs
	SET title = ?, slug=?, content=?, updated_at=? 
	WHERE id = ?`

	// blog.UpdatedAt = time.Now()
	updatedTimeinit := time.Now()
	newTime := updatedTimeinit.Format("2006-01-02 15:04:05")

	// blog.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", updatedAtStr)

	blog.Slug = utils.GenSlug(blog.Title)

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Slug, blog.Content, newTime, blog.ID)
	return err
}
