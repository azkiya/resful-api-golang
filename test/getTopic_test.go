package test

import (
	"net/http"
	"net/http/httptest"
	"newsapp/database"
	"newsapp/handlers"
	"newsapp/models"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetTopic(t *testing.T) {
	client := &database.MockTopicClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			expectedCode: 200,
		},
		"should return 404": {
			id:           "",
			expectedCode: 404,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Get", test.id).Return(models.Topic{}, nil)
			}
			req, _ := http.NewRequest("GET", "/topics/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("/topics/:id", handlers.GetTopic(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Get")
			}
		})
	}
}
