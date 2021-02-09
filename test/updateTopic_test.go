package test

import (
	"net/http"
	"net/http/httptest"
	"newsapp/database"
	"newsapp/handlers"
	"newsapp/models"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateTopic(t *testing.T) {
	client := &database.MockTopicClient{}
	id := primitive.NewObjectID().Hex()

	tests := map[string]struct {
		id           string
		payload      string
		expectedCode int
	}{
		"should return 200": {
			id:           id,
			payload:      `{"name":"Investment"}`,
			expectedCode: 200,
		},
		"should return 400": {
			id:           "",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.expectedCode == 200 {
				client.On("Update", test.id, mock.Anything).Return(models.TopicUpdate{}, nil)
			}

			req, _ := http.NewRequest("PATCH", "/topics/"+test.id, strings.NewReader(test.payload))
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.PATCH("topics/:id", handlers.UpdateTopic(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Update")
			}
		})
	}
}
