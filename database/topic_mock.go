package database

import (
	"newsapp/models"

	"github.com/stretchr/testify/mock"
)

type MockTopicClient struct {
	mock.Mock
}

func (m *MockTopicClient) Insert(topic models.Topic) (models.Topic, error) {
	args := m.Called(topic)
	return args.Get(0).(models.Topic), args.Error(1)
}

func (m *MockTopicClient) Update(id string, update interface{}) (models.TopicUpdate, error) {
	args := m.Called(id, update)
	return args.Get(0).(models.TopicUpdate), args.Error(1)
}

func (m *MockTopicClient) Delete(id string) (models.TopicDelete, error) {
	args := m.Called(id)
	return args.Get(0).(models.TopicDelete), args.Error(1)
}

func (m *MockTopicClient) Get(id string) (models.Topic, error) {
	args := m.Called(id)
	return args.Get(0).(models.Topic), args.Error(1)
}

func (m *MockTopicClient) Search(filter interface{}) ([]models.Topic, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Topic), args.Error(1)
}
