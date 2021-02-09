package handlers

import (
	"net/http"
	"newsapp/database"
	"newsapp/models"

	"github.com/gin-gonic/gin"
)

func InsertNews(db database.NewsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		news := models.News{}
		err := c.BindJSON(&news)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(news)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
