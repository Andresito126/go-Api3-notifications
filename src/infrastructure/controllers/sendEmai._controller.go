package controllers

import (
	"fmt"
	"net/http"

	"github.com/Andresito126/api3-notifications/src/application/services"
	"github.com/Andresito126/api3-notifications/src/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

// controlador para enviar correos
type SendEmailController struct {
	notification *services.Notification
}

// crea una nueva instancia del controlador
func NewSendEmailController() *SendEmailController {
	// obtener las dependencias
	resend := dependencies.GetResend()
	notification := services.NewNotification(resend)

	// inicializa el controlador
	return &SendEmailController{notification: notification}
}

//solicitud para enviar el correo
func (nse_c *SendEmailController) Run(ctx *gin.Context) {
	var data struct {
		StudentID string    `json:"student_id"`
		CourseID  string    `json:"course_id"`
		Status    string `json:"status"`
	}


	
	fmt.Println("Iniciando procesamiento de la solicitud para enviar correo...")

	
	if err := ctx.ShouldBindJSON(&data); err != nil {
		fmt.Printf("Error al parsear el JSON de la solicitud: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	
	fmt.Printf("Datos recibidos: StudentID: %d, CourseID: %d, Status: %s\n", data.StudentID, data.CourseID, data.Status)

	// valida que los datos sean correctos
	if data.Status == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "El estado de la inscripción es necesario"})
		return
	}

	// mensaje gmail
	message := fmt.Sprintf(
		"Estimado estudiante,\n\nTu inscripción en el curso con ID %d ha sido %s.\n\nSaludos,\nUniversidad Politécnica de Chiapas",
		data.CourseID, data.Status,
	)

	
	fmt.Println("Enviando correo con el siguiente contenido:", message)

	// servicio para enviar el correo
	nse_c.notification.Run(message)

	
	ctx.JSON(http.StatusOK, gin.H{"Message": "Correo enviado exitosamente"})
}
