package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware logs request details
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record the start time
		start := time.Now()

		// Log the request method and URL
		fmt.Printf("Started %s %s\n", c.Request.Method, c.Request.URL.Path)

		// Pass on to the next handler
		c.Next()

		// Log the time taken to process the request
		fmt.Printf("Completed %s in %v\n", c.Request.URL.Path, time.Since(start))
	}
}
