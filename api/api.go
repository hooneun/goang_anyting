package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hooneun/golang_anyting/model"
)

// HandlerInterface !
type HandlerInterface interface {
	GetUserByID(c *gin.Context)
}

// API !
func API() (*Handler, error) {
	db, err := model.NewORM()
	if err != nil {
		return nil, err
	}

	return &Handler{
		db: db,
	}, nil
}

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

// GetUserByID !
func (h *Handler) GetUserByID(c *gin.Context) {
	pID := c.Param("id")
	id, err := strconv.Atoi(pID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.db.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}
