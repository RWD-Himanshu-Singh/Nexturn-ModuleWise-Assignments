package controller

import (
	"blog_management/models"
	"blog_management/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type BlogController struct {
	BlogService *services.BlogService
}

func (c *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdBlog, err := c.BlogService.CreateBlog(&blog)
	if err != nil {
		http.Error(w, "Error creating blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBlog)
}

func (c *BlogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	blog, err := c.BlogService.GetBlog(id)
	if err != nil {
		http.Error(w, "Blog not found: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (c *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	blogs, err := c.BlogService.GetAllBlogs()
	if err != nil {
		http.Error(w, "Error fetching blogs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func (c *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var blog models.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if blog.ID != 0 && blog.ID != id {
		http.Error(w, "Mismatched blog ID", http.StatusBadRequest)
		return
	}

	blog.ID = id

	updatedBlog, err := c.BlogService.UpdateBlog(&blog)
	if err != nil {
		http.Error(w, "Error updating blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBlog)
}

func (c *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = c.BlogService.DeleteBlog(id)
	if err != nil {
		http.Error(w, "Error deleting blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Blog deleted successfully"})
}
