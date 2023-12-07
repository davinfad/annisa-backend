package repository

import (
	"annisa-salon/models"
	"database/sql"
	"time"
)

type BlogRepository interface {
	CreateBlog(blog *models.Blog) (*models.Blog, error)
    FindByBlogSlug(slug string) (*models.Blog, error)
	FindAllBlog(page, limit int) ([]*models.Blog, error)
    UpdateBlog(blog *models.Blog) (*models.Blog, error)
    DeleteBlog(slug string) (error)
}

type blogRepository struct {
	db *sql.DB
}

func NewBlogRepository (db *sql.DB) *blogRepository{
	return &blogRepository{db}
}

func (r *blogRepository) DeleteBlog (slug string) (error){
    query := "DELETE FROM blogs WHERE Slug = ?"

    _ , err := r.db.Exec(query, slug)
    if err != nil {
        return err
    }

    return nil
}

func (r *blogRepository) UpdateBlog(blog *models.Blog) (*models.Blog, error) {
    query := "UPDATE blogs SET Title=?, Description=?, FileName=?, UpdatedAt=? WHERE Slug=?"
    currentTime := time.Now().Format("2006-01-02 15:04:05")

    _, err := r.db.Exec(
        query,
        blog.Title,
        blog.Description,
        blog.FileName,
        currentTime,
        blog.Slug,
    )

    if err != nil {
        return nil, err
    }

    return blog, nil
}

func (r *blogRepository) CreateBlog(blog *models.Blog) (*models.Blog, error) {
    // currentTime := time.Now() // Mendapatkan waktu saat ini
	query := "INSERT INTO blogs (Title, Description, FileName, Slug ) VALUES (?, ?, ?, ?)"
    result, err := r.db.Exec(query, blog.Title, blog.Description, blog.FileName, blog.Slug)
    if err != nil {
        return blog, err
    }

    blogID, _ := result.LastInsertId()
    blog.ID = int(blogID)
    // blog.CreatedAt = currentTime

    return blog, nil
}

func (r *blogRepository) FindByBlogSlug(slug string) (*models.Blog, error) {
    query := "SELECT ID, Title, Description, FileName, Slug FROM blogs WHERE Slug = ?"

    row := r.db.QueryRow(query, slug)

    blog := &models.Blog{}
    err := row.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.FileName, &blog.Slug)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // Blog tidak ditemukan
        }
        return nil, err // Kesalahan lainnya
    }

    return blog, nil
}

func (r *blogRepository) FindAllBlog(page, limit int) ([]*models.Blog, error) {
    offset := (page - 1) * limit

    query := "SELECT ID, Title, Description, FileName, Slug FROM blogs LIMIT ?, ?"
    rows, err := r.db.Query(query, offset, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var blogs []*models.Blog
    for rows.Next() {
        blog := &models.Blog{}
        err := rows.Scan(&blog.ID, &blog.Title, &blog.Description, &blog.FileName, &blog.Slug)
        if err != nil {
            return nil, err
        }
        blogs = append(blogs, blog)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return blogs, nil
}


