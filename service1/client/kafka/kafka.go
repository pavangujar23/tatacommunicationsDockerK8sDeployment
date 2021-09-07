package kafka

import (
	"log"
)

type KafkaClient interface {
	Publish(string) error
}

type kafkaClient struct{}

func GetNewKafka() KafkaClient {
	return &kafkaClient{}
}

// publish data to consumer
func (k *kafkaClient) Publish(val string) error {
	_, err := producer.Write([]byte(val))
	if err != nil {
		log.Println("Error while writing data to the partition")
		panic(err)
	}
	log.Printf("Data sent : %s \n", val)
	return nil
}
