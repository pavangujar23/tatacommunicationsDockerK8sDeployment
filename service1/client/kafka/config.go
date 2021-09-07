package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/service1/config"
)

var (
	producer *kafka.Conn

	topic     = config.Conf.Topic
	partition = config.Conf.Partition
	kafcaport = config.Conf.KafKaPort
)

func init() {

	var err error
	producer, err = kafka.DialLeader(context.Background(), "tcp", kafcaport, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
}
