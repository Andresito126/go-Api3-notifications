package adapters

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/resend/resend-go/v2"
)

type Resend struct {}

func NewResend() *Resend {
	return &Resend{}
}

func (r *Resend) SendEmailToStudent(message string) {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	apiKey := os.Getenv("API_KEY")

	client := resend.NewClient(apiKey)

	mensaje := fmt.Sprintf("<strong>%s</strong>", message)
	fmt.Print(mensaje)

	params := &resend.SendEmailRequest{
        From:    "Acme <onboarding@resend.dev>",
        To:      []string{"andreju1260@gmail.com"},
        Html:     mensaje,
        Subject: "Hello from Golang",
        Cc:      []string{"cc@example.com"},
        Bcc:     []string{"bcc@example.com"},
        ReplyTo: "replyto@example.com",
    }

    sent, err := client.Emails.Send(params)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(sent.Id)

}