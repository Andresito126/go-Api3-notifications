package request

import (
	"bytes"
	"log"
	"net/http"
)

// fetch  para procesar el correo
func Fetch(jsonPayload []byte) {
	
	url := "http://localhost:8082/notification"

	// petición
	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Println("Error al hacer la petición:", err)
		return
	}

	// respuesta verificacion
	if res.StatusCode != http.StatusOK {
		log.Printf("Error al mandar el mensaje, código: %d", res.StatusCode)
		return
	}

	log.Println("Notificación procesada exitosamente")
}
