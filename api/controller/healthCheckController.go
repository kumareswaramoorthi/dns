package controller

import (
	"net/http"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/kumareswaramoorthi/dns/api/constants"
	"github.com/kumareswaramoorthi/dns/api/logging"
	"github.com/kumareswaramoorthi/dns/api/models"
)

type HealthCheckController interface {
	HealthCheck(c *gin.Context)
}

type healthCheckController struct {
}

func NewHealthCheckController() HealthCheckController {
	return healthCheckController{}
}

func (ctrl healthCheckController) HealthCheck(c *gin.Context) {
	logger := logging.GetLogger(c).
		WithField(constants.ReqID, requestid.Get(c)).
		WithField(constants.Interface, "LocationController").
		WithField(constants.Method, "GetLocation")

	logger.Info("Health check method Initiated")
	c.JSON(http.StatusOK, models.HealthResponse{
		Status:  "UP",
		Version: constants.AppVersion,
	})
	logger.Info("HealthCheck call completed")
}
