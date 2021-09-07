package app

import (
	"github.com/gin-gonic/gin"
	"github.com/service1/client/kafka"
	"github.com/service1/config"
	"github.com/service1/controller"
	"github.com/service1/service"
)

var (
	router           = gin.Default()
	applicattionPort = config.Conf.ApplicattionPort
)

// StartApplication starts gin application in default mode
func StartApplication() {

	var kafcaClient = kafka.GetNewKafka()
	var parkingService = service.GetNewParkingService(kafcaClient)
	var parkingController = controller.GetNewParkingController(parkingService)

	router.POST("/fillSlot", parkingController.FillSlot)
	router.POST("/emptySlot", parkingController.EmptySlot)

	router.Run(applicattionPort)
}
