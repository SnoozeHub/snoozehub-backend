//go:build prod

package dev_vs_prod

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

func IsAuthorized(publicKey string) bool {
	for _, k := range []string{
		"0x25072e991a78d25Ccc9DBEB1C4e787D19785F5D8",
		"0xe959e3f505694c35cee47e46E117E20d1Fa74191",
		"0x9D03FFa73F780a9A8760F9A0297E89663f9Bc0C2",
		"0xa5Bcf80D7dd0fF05031a6986B61aCd541E53201D",
		"0x283d0bE91d20f3D28142E8dBE69a61b0e46AF555",
	} {
		if k == publicKey {
			return true
		}
	}
	return false
}
