package controller

import (
	"encoding/json"
	"fmt"

	"github.com/service2/models"
	"github.com/service2/service"
)

func FillSlot(req string) error {

	var request models.FillSlotRequest

	//unmarshall json to models.FillSlotRequest
	err := json.Unmarshal([]byte(req), &request)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling Json")
	}

	//call service to do put call
	err = service.FillSlot(request)
	if err != nil {
		return err
	}
	return nil
}

func EmptySlot(req string) error {

	var request models.EmptySlotRequest

	//unmarshall json to models.FillSlotRequest
	err := json.Unmarshal([]byte(req), &request)
	if err != nil {
		return fmt.Errorf("Error while unmarshaling Json")
	}

	//call service to do put call
	err = service.EmptySlot(request)
	if err != nil {
		return err
	}
	return nil
}
