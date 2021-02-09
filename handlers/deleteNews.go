package handlers

import (
	"net/http"
	"newsapp/database"

	"github.com/gin-gonic/gin"
)

func DeleteNews(db database.NewsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		res, err := db.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
