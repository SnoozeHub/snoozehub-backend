//go:build !prod

package mail

import "log"

func Init() {
	
}

func Send(to string, subject string, message string) error{
	log.Println("mail to: " + to + ", subject: " + subject + ", message: " + message)
	return nil
}