package main

import (
	"log"

	"github.com/Andresito126/api3-notifications-consumer/src/rabbit"
)

func main() {
	//  instancia de rabbit
	rabbitClient := rabbit.NewRabbit()

	// recibir mensajes de la cola
	msgs := rabbitClient.ReceiveContent()

	// procesa los mensajes
	rabbit.ProcessMessage(msgs)

	log.Println("Consumer is running...")
}
