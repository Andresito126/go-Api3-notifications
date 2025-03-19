package repositories

type IResend interface {
    SendEmailToStudent(message string) 
}