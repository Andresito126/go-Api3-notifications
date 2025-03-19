package services

import (
	"github.com/Andresito126/api3-notifications/src/application/repositories"
)

type Notification struct {
	resend repositories.IResend
}

func NewNotification(resend repositories.IResend) *Notification {
	return &Notification{resend: resend}
}

// para enviar el correo
func (n *Notification) Run(message string) {
	// envia el correo con el adapter
	n.resend.SendEmailToStudent(message)
}