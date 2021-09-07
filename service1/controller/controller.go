package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/service1/models"
	"github.com/service1/service"
)

type ParkingController interface {
	FillSlot(c *gin.Context)
	EmptySlot(c *gin.Context)
}

type parkingController struct {
	parkingService service.ParkingService
}

func GetNewParkingController(parkingService service.ParkingService) ParkingController {
	return &parkingController{parkingService: parkingService}
}

func (controller *parkingController) FillSlot(c *gin.Context) {

	var request models.FillSlotRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.parkingService.FillSlot(request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Vehicle succesfully alloted slot : "+request.SpotId)
}

func (controller *parkingController) EmptySlot(c *gin.Context) {
	var request models.EmptySlotRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = controller.parkingService.EmptySlot(request)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Slot %s was succesfully emptied", request.SpotId))

}
