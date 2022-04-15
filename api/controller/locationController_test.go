package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kumareswaramoorthi/dns/api/controller/mocks"
	"github.com/kumareswaramoorthi/dns/api/errors"
	"github.com/kumareswaramoorthi/dns/api/models"
	"github.com/stretchr/testify/suite"
)

type LocationControllerTestSuite struct {
	suite.Suite
	context             *gin.Context
	recorder            *httptest.ResponseRecorder
	mockCtrl            *gomock.Controller
	mockLocationService *mocks.MockLocationService
	locationController  LocationController
}

func TestLocationController(t *testing.T) {
	suite.Run(t, new(LocationControllerTestSuite))
}

func (suite LocationControllerTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *LocationControllerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.mockLocationService = mocks.NewMockLocationService(suite.mockCtrl)
	suite.locationController = NewLocationController(suite.mockLocationService)
}

func (suite *LocationControllerTestSuite) TestGetLocationSuccessfully() {

	payload := models.Coordinates{
		X:   "123.12",
		Y:   "456.56",
		Z:   "789.89",
		Vel: "20.0",
	}

	req, _ := json.Marshal(payload)
	expectedResponse := &models.Location{Loc: 1389.57}
	response, _ := json.Marshal(expectedResponse)
	suite.context.Request, _ = http.NewRequest("POST", "/location", bytes.NewBufferString(string(req)))
	suite.mockLocationService.EXPECT().GetLocation(suite.context, payload).Return(expectedResponse, nil)
	suite.locationController.GetLocation(suite.context)
	suite.Equal(http.StatusOK, suite.recorder.Code)
	suite.JSONEq(string(response), suite.recorder.Body.String())
}

func (suite *LocationControllerTestSuite) TestGetLocationFailsIfNoPayload() {
	suite.context.Request, _ = http.NewRequest("POST", "/location", nil)
	suite.locationController.GetLocation(suite.context)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
}

func (suite *LocationControllerTestSuite) TestGetLocationFailsIfPayloadInvalid() {

	payload := models.Coordinates{
		X:   "abc",
		Y:   "456.56",
		Z:   "789.89",
		Vel: "20.0",
	}

	req, _ := json.Marshal(payload)
	suite.context.Request, _ = http.NewRequest("POST", "/location", bytes.NewBufferString(string(req)))
	suite.mockLocationService.EXPECT().GetLocation(suite.context, payload).Return(nil, errors.ErrBadRequest)
	suite.locationController.GetLocation(suite.context)
	suite.Equal(http.StatusBadRequest, suite.recorder.Code)
}
