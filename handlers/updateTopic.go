package handlers

import (
	"net/http"
	"newsapp/database"

	"github.com/gin-gonic/gin"
)

func UpdateTopic(db database.TopicInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var topic interface{}
		id := c.Param("id")
		err := c.BindJSON(&topic)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := db.Update(id, topic)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
