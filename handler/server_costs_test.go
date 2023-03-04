package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lmaraite/hcloud-cost-service/controller"
	"github.com/lmaraite/hcloud-cost-service/handler"
	"github.com/stretchr/testify/assert"
)

type ServerCostsControllerMock struct {
	response *controller.ServerCostsResponse
	err      error
}

func (mock *ServerCostsControllerMock) CalculateServerCosts() (*controller.ServerCostsResponse, error) {
	return mock.response, mock.err
}

func TestServerCostsHandler(t *testing.T) {
	testCases := []struct {
		name         string
		response     *controller.ServerCostsResponse
		err          error
		expectedCode int
	}{
		{
			name: "TestBasicRequest",
			response: &controller.ServerCostsResponse{
				Monthly: 13.00,
				Hourly:  0.012,
			},
			err:          nil,
			expectedCode: 200,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockController := &ServerCostsControllerMock{
				response: tc.response,
				err:      tc.err,
			}
			router := gin.Default()
			router.GET("/server/costs", func(ctx *gin.Context) {
				handler := handler.NewServerCostsHandler(mockController)
				handler.ServerCosts(ctx)
			})

			recorder := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", "/server/costs", nil)
			req.Header.Add("API_TOKEN", "12345")

			router.ServeHTTP(recorder, req)

			var response *controller.ServerCostsResponse
			json.Unmarshal(recorder.Body.Bytes(), &response)

			assert.Equal(t, tc.expectedCode, recorder.Code)
			assert.Equal(t, tc.response, response)

		})
	}
}
