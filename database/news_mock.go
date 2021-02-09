package database

import (
	"newsapp/models"

	"github.com/stretchr/testify/mock"
)

type MockNewsClient struct {
	mock.Mock
}

func (m *MockNewsClient) Insert(news models.News) (models.News, error) {
	args := m.Called(news)
	return args.Get(0).(models.News), args.Error(1)
}

func (m *MockNewsClient) Update(id string, update interface{}) (models.NewsUpdate, error) {
	args := m.Called(id, update)
	return args.Get(0).(models.NewsUpdate), args.Error(1)
}

func (m *MockNewsClient) Delete(id string) (models.NewsDelete, error) {
	args := m.Called(id)
	return args.Get(0).(models.NewsDelete), args.Error(1)
}

func (m *MockNewsClient) Get(id string) (models.News, error) {
	args := m.Called(id)
	return args.Get(0).(models.News), args.Error(1)
}

func (m *MockNewsClient) Search(filter interface{}) ([]models.News, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.News), args.Error(1)
}
