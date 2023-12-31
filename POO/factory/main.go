package main

import "fmt"

/*
**************************************************
Problema: Enviar notificaciones de Email y SMS.

Métodos/Formas de enviar la notificación:
- Email
- SMS

Canales/Tecnologías para enviar la notificación:
- Twilio (Para SMS)
- SES (Para Email)
**************************************************
*/

// Interfaz encargada enviar una notificación, según un método y un canal de envío.
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

// Interfaz que recoge el método y canal de envío.
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

/*
**********************************
	SMS
**********************************
*/
// STRUCT de SMS
type SMSNotification struct {
}

// Métodos de SMSNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// STRUCT para el envío de notificaciones SMS
type SMSNotificationSender struct {
}

// Métodos de SMSNotificationSender
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

/*
**********************************
	EMAIL
**********************************
*/
// STRUCT de Email
type EmailNotification struct {
}

// Métodos de EmailNotification
func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// STRUCT para el envío de notificaciones Email
type EmailNotificationSender struct {
}

// Métodos de EmailNotificationSender
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

/*
**********************************
	Funcionalidad
**********************************
*/
// Función que evalua el tipo de notificación que debe enviar
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("no notification type")
}

// Función que envía la notificación
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

// Función que escoge la forma/método por la que se enviará la notificación
func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

// Función que escoge el canal/tecnología por la que se enviará la notificación
func getChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	getMethod(smsFactory)        // Recoge el método de Envío
	getChannel(smsFactory)       // Recoge el canal de Envío
	sendNotification(smsFactory) // Envía la notificación

	getMethod(emailFactory)        // Recoge el método de Envío
	getChannel(emailFactory)       // Recoge el canal de Envío
	sendNotification(emailFactory) // Envía la notificación
}
