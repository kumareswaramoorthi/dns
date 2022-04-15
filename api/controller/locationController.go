package controller

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/kumareswaramoorthi/dns/api/constants"
	"github.com/kumareswaramoorthi/dns/api/errors"
	"github.com/kumareswaramoorthi/dns/api/logging"
	"github.com/kumareswaramoorthi/dns/api/models"
	"github.com/kumareswaramoorthi/dns/api/service"
)

type LocationController interface {
	GetLocation(c *gin.Context)
}

type locationController struct {
	locationService service.LocationService
}

func NewLocationController(locationService service.LocationService) LocationController {
	return locationController{
		locationService: locationService,
	}
}

// @title Drone Navigation Service API
// Get Location godoc
// @Tags Get Location
// @Accept json
// @Produce  json
// @Description Get Location
// @Success 200 {object} models.Location
// @Failure 400 {object} errors.ErrorResponse
// @Param Coordinates body models.Coordinates true "request body"
// @Router /location [POST]
func (ctrl locationController) GetLocation(c *gin.Context) {
	logger := logging.GetLogger(c).
		WithField(constants.ReqID, requestid.Get(c)).
		WithField(constants.Interface, "LocationController").
		WithField(constants.Method, "GetLocation")
	logger.Info("GetLocation method Initiated")

	var coordinates models.Coordinates

	//Bind json to coordinates object
	if err := c.ShouldBindJSON(&coordinates); err != nil {
		logger.Errorf("ShouldBindJSON - %s", err.Error())
		c.AbortWithStatusJSON(errors.ErrBadRequest.HttpStatusCode, errors.ErrBadRequest)
		return
	}

	//find source and destination
	loc, err := ctrl.locationService.GetLocation(c, coordinates)
	if err != nil {
		logger.Errorf("GetLocation - %s", err.Error())
		c.AbortWithStatusJSON(err.HttpStatusCode, err)
		return
	}

	//return the response
	c.JSON(http.StatusOK, loc)
	logger.Info("GetLocation call completed")
}
