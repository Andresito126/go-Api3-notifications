package dependencies

import (
	"github.com/Andresito126/api3-notifications/src/infrastructure/adapters"
)

var resend *adapters.Resend

//
func InitDependencies() {
	resend = adapters.NewResend()
}


func GetResend() *adapters.Resend {
	return resend
}
