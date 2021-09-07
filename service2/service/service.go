package service

import (
	"fmt"
	"time"

	"github.com/service2/models"
	"github.com/service2/models/repo"
)

func GetEmptyParkingSpots() (parkingList []repo.Parking, err error) {
	result, err := repo.GetEmptyParkingSpots()
	return result, err
}

func GetAllParkingSpots() (parkingList []repo.Parking, err error) {
	result, err := repo.GetAllParkingSpots()
	return result, err
}

func FillSlot(request models.FillSlotRequest) error {

	parking, err := repo.GetParkingByID(request.SpotId)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if parking.Occupied == "true" {
		return fmt.Errorf("Parking Spot Already occupied")
	} else {
		p := repo.Parking{SpotId: request.SpotId, VehicleId: request.VehicleId, Occupied: "true", StartTime: time.Now().Unix(), EndTime: 0}
		err = repo.PutParkingSpot(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func EmptySlot(request models.EmptySlotRequest) error {

	parking, err := repo.GetParkingByID(request.SpotId)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if parking.Occupied == "false" {
		return fmt.Errorf("No vehicle is parked in spot : " + request.SpotId)
	} else {
		p := repo.Parking{SpotId: request.SpotId, VehicleId: "", Occupied: "false", StartTime: 0, EndTime: 0}
		err = repo.PutParkingSpot(p)
		if err != nil {
			return err
		}

		//call billing service

		parking.EndTime = time.Now().Unix()
		fmt.Println("calling billing service with data : ", parking)
	}
	return nil
}
