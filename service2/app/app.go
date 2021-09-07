package app

import (
	"github.com/gin-gonic/gin"
	"github.com/service2/config"
	"github.com/service2/controller"
)

var (
	router           = gin.Default()
	applicattionPort = config.Conf.ApplicattionPort
)

// StartApplication starts gin application
func StartApplication() {

	router.GET("/GetEmptyParkingSpots", controller.GetEmptyParkingSpots)
	router.GET("/GetAllParkingSpots", controller.GetAllParkingSpots)

	router.Run(applicattionPort)
}
