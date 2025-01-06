package repository

import (
	"blog_management/models"
	"database/sql"
	"log"
	"time"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

func (r *BlogRepository) CreateBlog(blog *models.Blog) (*models.Blog, error) {
	stmt, err := r.DB.Prepare("INSERT INTO blogs (title, content, author, timestamp) VALUES (?, ?, ?, CURRENT_TIMESTAMP)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Content, blog.Author)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	blog.ID = int(id)
	blog.Timestamp = time.Now()
	return blog, nil
}

func (repo *BlogRepository) GetBlog(id int) (*models.Blog, error) {
	row := repo.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id)
	blog := &models.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepository) GetAllBlogs() ([]models.Blog, error) {
	rows, err := repo.DB.Query("SELECT id, title, content, author, timestamp FROM blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repo *BlogRepository) UpdateBlog(blog *models.Blog) (*models.Blog, error) {
	stmt, err := repo.DB.Prepare("UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, blog.ID)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (repo *BlogRepository) DeleteBlog(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
