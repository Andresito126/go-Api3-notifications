package rabbit

import (
	"fmt"
	"log"

	"github.com/Andresito126/api3-notifications-consumer/src/request"
	amqp "github.com/rabbitmq/amqp091-go"
)

// procesa el mensaje
func ProcessMessage(msgs <-chan amqp.Delivery) {
	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("[x] Mensaje recibido: %s", d.Body)

			fmt.Println("Mensaje recibido: ", d.Body)

			// llama a la funcion fetch de la solicitud a la api3
			request.Fetch(d.Body)
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
