package database

import (
	"context"
	"encoding/json"
	"newsapp/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsInterface interface {
	Insert(models.News) (models.News, error)
	Update(string, interface{}) (models.NewsUpdate, error)
	Delete(string) (models.NewsDelete, error)
	Get(string) (models.News, error)
	Search(interface{}) ([]models.News, error)
}

type NewsClient struct {
	Ctx context.Context
	Col *mongo.Collection
}

func (c *NewsClient) Insert(docs models.News) (models.News, error) {
	News := models.News{}

	res, err := c.Col.InsertOne(c.Ctx, docs)
	if err != nil {
		return News, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.Get(id)
}

func (c *NewsClient) Update(id string, update interface{}) (models.NewsUpdate, error) {
	result := models.NewsUpdate{
		ModifiedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}

	News, err := c.Get(id)
	if err != nil {
		return result, err
	}

	var exist map[string]interface{}
	b, err := json.Marshal(News)
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

	NewsUpdated, err := c.Get(id)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = res.ModifiedCount
	result.Result = NewsUpdated
	return result, nil

}

func (c *NewsClient) Delete(id string) (models.NewsDelete, error) {
	result := models.NewsDelete{
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

func (c *NewsClient) Get(id string) (models.News, error) {
	News := models.News{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return News, err
	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&News)
	if err != nil {
		return News, err
	}

	return News, nil
}

func (c *NewsClient) Search(filter interface{}) ([]models.News, error) {
	Newss := []models.News{}
	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := c.Col.Find(c.Ctx, filter)
	if err != nil {
		return Newss, err
	}

	for cursor.Next(c.Ctx) {
		row := models.News{}
		cursor.Decode(&row)
		Newss = append(Newss, row)
	}

	return Newss, nil
}
