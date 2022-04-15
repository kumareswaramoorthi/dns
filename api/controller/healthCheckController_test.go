package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kumareswaramoorthi/dns/api/models"
	"github.com/stretchr/testify/suite"
)

type HealthCheckControllerTestSuite struct {
	suite.Suite
	context               *gin.Context
	recorder              *httptest.ResponseRecorder
	mockCtrl              *gomock.Controller
	healthCheckController HealthCheckController
}

func TestHealthCheckController(t *testing.T) {
	suite.Run(t, new(HealthCheckControllerTestSuite))
}

func (suite HealthCheckControllerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *HealthCheckControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.healthCheckController = NewHealthCheckController()
}

func (suite HealthCheckControllerTestSuite) TestHealthCheckSuccessfully() {
	context := suite.context
	context.Request, _ = http.NewRequest("GET", "/", nil)
	healthDetails := models.HealthResponse{
		Status:  "UP",
		Version: "1.0.0",
	}
	expectedResponse, _ := json.Marshal(healthDetails)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.healthCheckController.HealthCheck(context)
	suite.JSONEq(string(expectedResponse), suite.recorder.Body.String())
}
