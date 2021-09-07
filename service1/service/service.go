package service

import (
	"encoding/json"

	"github.com/service1/client/kafka"
	"github.com/service1/models"
)

type ParkingService interface {
	FillSlot(models.FillSlotRequest) error
	EmptySlot(models.EmptySlotRequest) error
}

type parkingService struct {
	kafcaClient kafka.KafkaClient
}

func GetNewParkingService(kafcaClient kafka.KafkaClient) ParkingService {
	return &parkingService{kafcaClient: kafcaClient}
}

func (service *parkingService) FillSlot(fillSlot models.FillSlotRequest) error {

	fillSlotbyte, _ := json.Marshal(fillSlot)

	err := service.kafcaClient.Publish("FillSlot" + "#" + string(fillSlotbyte))
	if err != nil {
		return err
	}

	return nil
}

func (service *parkingService) EmptySlot(emptySlot models.EmptySlotRequest) error {

	emptySlotbyte, _ := json.Marshal(emptySlot)

	err := service.kafcaClient.Publish("EmptySlot" + "#" + string(emptySlotbyte))
	if err != nil {
		return err
	}
	return nil
}
