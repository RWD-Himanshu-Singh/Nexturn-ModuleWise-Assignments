package services

import (
	"blog_management/models"
	"blog_management/repository"
)

type BlogService struct {
	BlogRepo *repository.BlogRepository
}

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{BlogRepo: repo}
}

func (s *BlogService) CreateBlog(blog *models.Blog) (*models.Blog, error) {
	return s.BlogRepo.CreateBlog(blog)
}

func (s *BlogService) GetBlog(id int) (*models.Blog, error) {
	return s.BlogRepo.GetBlog(id)
}

func (s *BlogService) GetAllBlogs() ([]models.Blog, error) {
	return s.BlogRepo.GetAllBlogs()
}

func (s *BlogService) UpdateBlog(blog *models.Blog) (*models.Blog, error) {
	return s.BlogRepo.UpdateBlog(blog)
}

func (s *BlogService) DeleteBlog(id int) error {
	return s.BlogRepo.DeleteBlog(id)
}
