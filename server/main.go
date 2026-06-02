package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Gortex backend")
	})

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}
