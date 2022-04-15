package service

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kumareswaramoorthi/dns/api/errors"
	"github.com/kumareswaramoorthi/dns/api/models"
	"github.com/stretchr/testify/suite"
)

type LocationServiceTestSuite struct {
	suite.Suite
	context         *gin.Context
	recorder        *httptest.ResponseRecorder
	mockCtrl        *gomock.Controller
	locationService LocationService
}

func TestLocationService(t *testing.T) {
	suite.Run(t, new(LocationServiceTestSuite))
}

func (suite LocationServiceTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *LocationServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.locationService = NewLocationService()
}

func (suite *LocationServiceTestSuite) TestGetLocationSuccessfully() {
	payload := models.Coordinates{
		X:   "123.12",
		Y:   "456.56",
		Z:   "789.89",
		Vel: "20.0",
	}

	actualResponse, err := suite.locationService.GetLocation(suite.context, payload)
	expectedResponse := &models.Location{Loc: 1389.57}
	suite.Equal(expectedResponse, actualResponse)
	suite.Nil(err)
}

func (suite *LocationServiceTestSuite) TestGetLocationFailIfPayloadInvalid() {
	xInvalidReq := models.Coordinates{X: "abc", Y: "12.34", Z: "45.67", Vel: "67.43"}
	yInvalidReq := models.Coordinates{X: "12.6", Y: "qbc", Z: "45.67", Vel: "67.43"}
	zInvalidReq := models.Coordinates{X: "45", Y: "12.34", Z: "qbc&", Vel: "67.43"}
	velInvalidReq := models.Coordinates{X: "12.12", Y: "12.34", Z: "45.67", Vel: "%^&"}

	var payload []models.Coordinates
	payload = append(payload, xInvalidReq, yInvalidReq, zInvalidReq, velInvalidReq)
	for _, v := range payload {
		_, err := suite.locationService.GetLocation(suite.context, v)
		suite.Equal(errors.ErrBadRequest, err)
		suite.NotNil(err)
	}
}
