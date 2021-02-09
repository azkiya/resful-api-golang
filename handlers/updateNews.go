package handlers

import (
	"net/http"
	"newsapp/database"

	"github.com/gin-gonic/gin"
)

func UpdateNews(db database.NewsInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var news interface{}
		id := c.Param("id")
		err := c.BindJSON(&news)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, news)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
