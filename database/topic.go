package database

import (
	"context"
	"encoding/json"
	"newsapp/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicInterface interface {
	Insert(models.Topic) (models.Topic, error)
	Update(string, interface{}) (models.TopicUpdate, error)
	Delete(string) (models.TopicDelete, error)
	Get(string) (models.Topic, error)
	Search(interface{}) ([]models.Topic, error)
}

type TopicClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *TopicClient) Insert(docs models.Topic) (models.Topic, error) {
	topic := models.Topic{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return topic, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *TopicClient) Update(id string, update interface{}) (models.TopicUpdate, error) {
	result := models.TopicUpdate{
		ModifiedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	topic, err := c.Get(id)
	if err != nil {
		return result, err
	}

	var exist map[string]interface{}
	b, err := json.Marshal(topic)
	if err != nil {
		return result, err
	}
	json.Unmarshal(b, &exist)

	change := update.(map[string]interface{})
	for key := range change {
		if change[key] == exist[key] {
			delete(change, key)
		}
	}

	if len(change) == 0 {
		return result, nil
	}

	res, err := c.Col.UpdateOne(c.Ctx, bson.M{"_id": _id}, bson.M{"$set": change})
	if err != nil {
		return result, err
	}

	topicUpdated, err := c.Get(id)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = res.ModifiedCount
	result.Result = topicUpdated
	return result, nil

}

func (c *TopicClient) Delete(id string) (models.TopicDelete, error) {
	result := models.TopicDelete{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	res, err := c.Col.DeleteOne(c.Ctx, bson.M{"_id": _id})
	if err != nil {
		return result, err
	}
	result.DeletedCount = res.DeletedCount
	return result, nil
}

func (c *TopicClient) Get(id string) (models.Topic, error) {
	topic := models.Topic{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return topic, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&topic)
	if err != nil {
		return topic, err
	}

	return topic, nil
}

func (c *TopicClient) Search(filter interface{}) ([]models.Topic, error) {
	topics := []models.Topic{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return topics, err
	}

	for cursor.Next(c.Ctx) {
		row := models.Topic{}
		cursor.Decode(&row)
		topics = append(topics, row)
	}

	return topics, nil
}
