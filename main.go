package main

import (
	"context"
	"newsapp/config"
	"newsapp/database"
	"newsapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &database.TopicClient{
		Col: collection,
		Ctx: ctx,
	}

	r := gin.Default()

	topics := r.Group("/topics")
	{
		topics.GET("/", handlers.SearchTopic(client))
		topics.GET("/:id", handlers.GetTopic(client))
		topics.POST("/", handlers.InsertTopic(client))
		topics.PATCH("/:id", handlers.UpdateTopic(client))
		topics.DELETE("/:id", handlers.DeleteTopic(client))

	}

	r.GET("/ping", handlers.Ping)

	r.Run(":8080")

}
