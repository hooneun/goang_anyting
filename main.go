package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tt/api"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/time", api.Time)
	r.GET("/user/:id", api.GetUserById)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
