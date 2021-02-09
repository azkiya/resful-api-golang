package test

import (
	"net/http"
	"net/http/httptest"
	"newsapp/database"
	"newsapp/handlers"
	"newsapp/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestSearchTopic(t *testing.T) {
	client := &database.MockTopicClient{}

	tests := map[string]struct {
		payload      string
		expectedCode int
	}{
		"should return 200": {
			payload:      `{"name":"Investment"}`,
			expectedCode: 200,
		},
		"should return 400": {
			payload:      "",
			expectedCode: 400,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client.On("Search", mock.Anything).Return([]models.Topic{}, nil)
			req, _ := http.NewRequest("GET", "/topics?q="+test.payload, nil)
			rec := httptest.NewRecorder()

			r := gin.Default()
			r.GET("topics", handlers.SearchTopic(client))
			r.ServeHTTP(rec, req)

			if test.expectedCode == 200 {
				client.AssertExpectations(t)
			} else {
				client.AssertNotCalled(t, "Search")
			}
		})
	}
}
