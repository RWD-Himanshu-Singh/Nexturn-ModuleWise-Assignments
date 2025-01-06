package main

import (
	"blog_management/config"
	"blog_management/controller"
	"blog_management/middleware" // Import the middleware package
	"blog_management/repository"
	"blog_management/services"
	"fmt"
	"net/http"
)

func main() {
	db, err := config.InitializeDatabase()
	if err != nil {
		fmt.Println("Error in Initializing Database: ", err)
	}
	defer db.Close()

	blogRepo := repository.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogController := controller.BlogController{BlogService: blogService}

	// Wrap routes with middleware
	http.Handle("/blogs", middleware.LoggingMiddleware(
		middleware.BasicAuthMiddleware(http.HandlerFunc(blogController.GetAllBlogs)),
	))

	http.Handle("/blog", middleware.LoggingMiddleware(
		middleware.BasicAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPost:
				blogController.CreateBlog(w, r)
			case http.MethodGet:
				blogController.GetBlog(w, r)
			case http.MethodPut:
				blogController.UpdateBlog(w, r)
			case http.MethodDelete:
				blogController.DeleteBlog(w, r)
			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
		})),
	))

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
