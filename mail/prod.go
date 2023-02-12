//go:build !prod

package mail

import (
	_ "embed"
	"io/ioutil"

	"github.com/mailgun/mailgun-go"
)

var mailgunSendindKey string

func Init() {
	tmp, _ := ioutil.ReadFile("secrets/mailgun-sending-key.key")
	mailgunSendindKey = string(tmp)
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
