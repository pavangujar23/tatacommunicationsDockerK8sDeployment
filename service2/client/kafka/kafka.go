package kafka

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
	"github.com/service2/config"
	"github.com/service2/controller"
)

var (
	kafcaport = config.Conf.KafKaPort
	topic     = config.Conf.Topic
	partition = config.Conf.Partition
)

// consume listens to message posted by the producer on topic
func Consume() {

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafcaport, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("failed to close connection:", err)
		}
	}()

	for {

		mes, err := conn.ReadMessage(100)
		if err != nil {
			continue
		}

		temp := strings.Split(string(mes.Value), "#")

		if len(temp) == 0 {
			fmt.Println("Invalid request Structure with :", string(mes.Value))
			continue
		}

		// temp[0] contains info on which controller to be called
		// temp[1] contains the json passed by the producer

		switch temp[0] {

		case "FillSlot":
			{
				err := controller.FillSlot(temp[1])
				if err != nil {
					fmt.Println(err.Error())
				}
			}

		case "EmptySlot":
			{
				err := controller.EmptySlot(temp[1])
				if err != nil {
					fmt.Println(err.Error())
				}
			}

		default:
			{
				fmt.Println("No required controller present")
			}

		}

	}
}
