//go:build !prod

package dev_vs_prod

import "log"

func Init() {

}

var LatestMessage string // used by tests

func Send(to string, subject string, message string) error {
	LatestMessage = message
	log.Println("mail to: " + to + ", subject: " + subject + ", message: " + message)
	return nil
}

func IsAuthorized(publicKey string) bool {
	return true
}
