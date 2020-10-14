package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hooneun/golang_anyting/api"
)

func setupRouter(h api.HandlerInterface) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/time", api.Time)
	r.GET("/user/:id", h.GetUserByID)

	return r
}

func main() {
	h, err := api.API()
	if err != nil {
		log.Fatal(err)
	}
	r := setupRouter(h)
	r.Run(":8080")
}
