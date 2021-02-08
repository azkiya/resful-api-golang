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
		topics.POST("/", handlers.InsertTopic(client))
	}

	topics.GET("/", handlers.InsertTopic(client))

	r.GET("/ping", handlers.Ping)

	r.Run(":8080")

}
