package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/service2/service"
)

func GetEmptyParkingSpots(c *gin.Context) {

	result, err := service.GetEmptyParkingSpots()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, result)
}

func GetAllParkingSpots(c *gin.Context) {

	result, err := service.GetAllParkingSpots()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, result)
}
