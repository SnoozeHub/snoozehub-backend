//go:build prod

package mail

import "log"

func Init() {
	
}

func Send(to string, message string) error{
	log.Println(to, " ", message)
	return nil
}