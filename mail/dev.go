//go:build !prod

package mail

import "log"

func Init() {
	
}

var LatestMessage string // used by tests

func Send(to string, subject string, message string) error{
	LatestMessage = message
	log.Println("mail to: " + to + ", subject: " + subject + ", message: " + message)
	return nil
}