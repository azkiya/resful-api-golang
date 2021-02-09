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

	clientNews := &database.NewsClient{
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

	news := r.Group("/news")
	{
		news.GET("/", handlers.SearchNews(clientNews))
		news.GET("/:id", handlers.GetNews(clientNews))
		news.POST("/", handlers.InsertNews(clientNews))
		news.PATCH("/:id", handlers.UpdateNews(clientNews))
		news.DELETE("/:id", handlers.DeleteNews(clientNews))

	}

	r.GET("/ping", handlers.Ping)

	r.Run(":8080")

}
