package main

import (
	"github.com/service2/app"
	"github.com/service2/client/kafka"
)

func main() {

	// kafka consumer listning to topic
	go kafka.Consume()

	//controller part of gin application
	app.StartApplication()

}
