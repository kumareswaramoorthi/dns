package router

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/kumareswaramoorthi/dns/api/controller"
	"github.com/kumareswaramoorthi/dns/api/logging"
	"github.com/kumareswaramoorthi/dns/api/middlewares"
	"github.com/kumareswaramoorthi/dns/api/service"
	"github.com/kumareswaramoorthi/dns/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
)

func SetupRouter() *gin.Engine {

	//Get default router from gin
	router := gin.Default()

	//create a global logger for the server
	apiLoggerEntry := logging.NewLoggerEntry()

	//enable logging middleware
	router.Use(logging.LoggingMiddleware(apiLoggerEntry))

	//enable unique request id generator  middleware
	router.Use(requestid.New())

	//enable rate limiter middleware
	router.Use(mgin.NewMiddleware(middlewares.RateLimiterMiddleware()))

	//swagger init
	docs.SwaggerInfo.Title = "Drone Navigation Service"
	docs.SwaggerInfo.Description = "This lists down the endpoints that are part of Drone Navigation Service API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	healthCheckController := controller.NewHealthCheckController()
	locationService := service.NewLocationService()
	locationController := controller.NewLocationController(locationService)

	//routes
	router.GET("/", healthCheckController.HealthCheck)
	router.POST("/location", locationController.GetLocation)

	return router
}
