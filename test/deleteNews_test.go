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

func TestDeleteNews(t *testing.T) {
	client := &database.MockNewsClient{}
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
				client.On("Delete", test.id).Return(models.NewsDelete{}, nil)
			}
			req, _ := http.NewRequest("DELETE", "/news/"+test.id, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.DELETE("/news/:id", handlers.DeleteNews(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Delete")
			}
		})
	}
}
