//go:build !prod

package mail

import (
	_ "embed"
	"os"

	"github.com/mailgun/mailgun-go"
)

var mailgunSendindKey string

func Init() {
	file, _ := os.Open("secrets/mailgun-sending-key.key")
	defer file.Close()
	fi, _ := file.Stat()
	data := make([]byte, fi.Size())
	_, _ = file.Read(data)
	mailgunSendindKey = string(data)
}

func Send(to string, subject string, message string) error {
	mg := mailgun.NewMailgun("etern.fun", mailgunSendindKey)
	mg.SetAPIBase("https://api.eu.mailgun.net/v3")
	m := mg.NewMessage(
		"Snoozehub <snoozehub@etern.fun>",
		subject,
		message,
		to,
	)
	_, _, err := mg.Send(m)
	return err

}
