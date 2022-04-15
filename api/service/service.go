package service

import (
	"math"
	"strconv"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/kumareswaramoorthi/dns/api/constants"
	"github.com/kumareswaramoorthi/dns/api/errors"
	"github.com/kumareswaramoorthi/dns/api/logging"
	"github.com/kumareswaramoorthi/dns/api/models"
)

type LocationService interface {
	GetLocation(c *gin.Context, coordinates models.Coordinates) (*models.Location, *errors.ErrorResponse)
}

type locationService struct {
}

func NewLocationService() LocationService {
	return &locationService{}
}

func (fts *locationService) GetLocation(c *gin.Context, coordinates models.Coordinates) (*models.Location, *errors.ErrorResponse) {
	logger := logging.GetLogger(c).
		WithField(constants.ReqID, requestid.Get(c)).
		WithField(constants.Interface, "LocationService").
		WithField(constants.Method, "GetLocation")

	var (
		location     models.Location
		x, y, z, vel float64
		err          error
	)

	if x, err = strconv.ParseFloat(coordinates.X, 64); err != nil {
		logger.Errorf("ParseFloat - %s", err.Error())
		return nil, errors.ErrBadRequest
	}
	if y, err = strconv.ParseFloat(coordinates.Y, 64); err != nil {
		logger.Errorf("ParseFloat - %s", err.Error())
		return nil, errors.ErrBadRequest
	}
	if z, err = strconv.ParseFloat(coordinates.Z, 64); err != nil {
		logger.Errorf("ParseFloat - %s", err.Error())
		return nil, errors.ErrBadRequest
	}
	if vel, err = strconv.ParseFloat(coordinates.Vel, 64); err != nil {
		logger.Errorf("ParseFloat - %s", err.Error())
		return nil, errors.ErrBadRequest
	}

	location.Loc = math.Round(((x*constants.SectorID)+(y*constants.SectorID)+(z*constants.SectorID)+vel)*100) / 100
	logger.Infof("location - %f", location.Loc)
	return &location, nil
}
