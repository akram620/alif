package handler

import (
	"bytes"
	"github.com/akram620/alif/internal/models"
	"github.com/akram620/alif/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http/httptest"
	"testing"
)

func TestHandler_CreateEvent(t *testing.T) {
	type mockBehaviour func(s *mocks.Events, event *models.Event)

	testTable := []struct {
		name          string
		inputBody     string
		inputEvent    *models.Event
		mockBehaviour mockBehaviour
		expectedCode  int
		expectedBody  string
	}{
		{
			name: "OK",
			inputBody: `{
						  "orderType": "Test1",
						  "sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
						  "card": "4433**1409",
						  "eventDate": "2024-04-24 19:08:52.835626 +05:00",
						  "websiteUrl": "https://amazon.com"
						}`,
			inputEvent: &models.Event{
				OrderType:  "Test1",
				SessionID:  "29827525-06c9-4b1e-9d9b-7c4584e82f56",
				Card:       "4433**1409",
				EventDate:  "2024-04-24 19:08:52.835626 +05:00",
				WebsiteURL: "https://amazon.com",
			},
			mockBehaviour: func(s *mocks.Events, event *models.Event) {
				s.On("CreateEvent", event).Return(nil)
			},
			expectedCode: 200,
			expectedBody: `{"status":"success"}`,
		},

		// other test cases
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init deps
			eventsMock := mocks.NewEvents(t)
			testCase.mockBehaviour(eventsMock, testCase.inputEvent)

			handler := NewHandler(eventsMock)
			handler.InitRoutes()

			// Test server
			r := gin.New()
			r.POST("/api/v1/events", handler.CreateEvent)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/events", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedCode, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())
		})
	}
}
