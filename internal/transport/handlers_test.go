package transport

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/tahanasir/service-catalog/internal/models"
	gomock "go.uber.org/mock/gomock"
)

func TestGetService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockdbStorage(ctrl)

	id := "550e8400-e29b-41d4-a716-446655440000"
	expectedID, _ := uuid.Parse(id)
	expectedService := models.SingleService{
		ID:   expectedID,
		Name: "Test Service",
	}

	mockStorage.EXPECT().GetService(expectedID).Return(expectedService, nil)

	testCases := []struct {
		Name           string
		RequestMethod  string
		RequestURL     string
		ExpectedStatus int
		ExpectedBody   models.SingleService
	}{
		{
			Name:           "GetService_Success",
			RequestMethod:  "GET",
			RequestURL:     "/v1/services/" + id,
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   expectedService,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(tc.RequestMethod, tc.RequestURL, nil)
			w := httptest.NewRecorder()

			handler := GetService(mockStorage)

			r := chi.NewRouter()
			r.Get("/v1/services/{id}", handler)
			r.ServeHTTP(w, req)

			assert.Equal(t, tc.ExpectedStatus, w.Code)

			var service models.SingleService
			err := json.NewDecoder(w.Body).Decode(&service)
			if err != nil {
				t.Fatalf("Failed to decode JSON response: %v", err)
			}

			assert.Equal(t, tc.ExpectedBody, service)
		})
	}
}
