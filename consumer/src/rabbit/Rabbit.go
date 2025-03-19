package rabbit

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	Broker  *amqp.Connection
	Channel *amqp.Channel
}


func NewRabbit() *Rabbit {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}


	rabbitUrl := os.Getenv("RABBIT_URL")

	// conectar 
	conn, err := amqp.Dial(rabbitUrl)
	if err != nil {
		log.Fatal("Error al abrir una conexión hacia RabbitMQ:", err)
	}

	// abrir un canal
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir un canal:", err)
	}

	return &Rabbit{Broker: conn, Channel: ch}
}

// recibe mensajes desde la cola
func (r *Rabbit) ReceiveContent() <-chan amqp.Delivery {
	// declarar el exchange
	err := r.Channel.ExchangeDeclare(
		"notificationsExchange", // Nombre del exchange
		"direct",                // Tipo de exchange
		true,                    // Durable
		false,                   // Auto-deleted
		false,                   // Internal
		false,                   // No-wait
		nil,                     // Argumentos
	)
	r.FailOnError(err, "Failed to declare an exchange")

	// declarar la cola
	q, err := r.Channel.QueueDeclare(
		"notificationsQueue", // Nombre de la cola
		true,                 // Durable
		false,                // Delete when unused
		false,                // Exclusive
		false,                // No-wait
		nil,                  // Argumentos
	)
	r.FailOnError(err, "Failed to declare a queue")

	// vinculacion de la cola al exchange
	err = r.Channel.QueueBind(
		q.Name,                // Nombre de la cola
		"notification",        // Routing key
		"notificationsExchange", // Nombre del exchange
		false,
		nil,
	)
	r.FailOnError(err, "Failed to bind a queue")

	// para recibir mensajes
	msgs, err := r.Channel.Consume(
		q.Name, // Nombre de la cola
		"",     // Consumer
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Args
	)
	r.FailOnError(err, "Failed to register a consumer")

	return msgs
}


func (r *Rabbit) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
