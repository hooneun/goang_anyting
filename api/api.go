package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hooneun/golang_anyting/model"
)

func Time(c *gin.Context) {
	db, err := model.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var time string
	err = db.QueryRow("SELECT now()").Scan(&time)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, map[string]string{
		"time": time,
	})
}

func GetUserById(c *gin.Context) {
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := model.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
