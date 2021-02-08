package handlers

import (
	"net/http"
	"newsapp/database"
	"newsapp/models"

	"github.com/gin-gonic/gin"
)

func InsertTopic(db database.TopicInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		topic := models.Topic{}
		err := c.BindJSON(&topic)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Insert(topic)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
