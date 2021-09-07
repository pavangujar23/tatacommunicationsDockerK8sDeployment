package models

type FillSlotRequest struct {
	SpotId    string `json:"spotId"`
	VehicleId string `json:"vehicleId"`
}

type EmptySlotRequest struct {
	SpotId string `json:"spotId"`
}
