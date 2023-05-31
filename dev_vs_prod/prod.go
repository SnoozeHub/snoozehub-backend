//go:build prod

package dev_vs_prod

import (
	_ "embed"
	"os"
	"strings"
	"github.com/mailgun/mailgun-go"
)

var mailgunSendindKey string

var publicKeyWhitelist []string

func Init() {
	file, _ := os.Open("secrets/mailgun-sending-key.key")
	fi, _ := file.Stat()
	data := make([]byte, fi.Size())
	_, _ = file.Read(data)
	mailgunSendindKey = string(data)
	file.Close()

	file, _ = os.Open("secrets/whitelist.txt")
	defer file.Close()
	fi, _ = file.Stat()
	data = make([]byte, fi.Size())
	_, _ = file.Read(data)
	lines := strings.Split(string(data), "\n")
 publicKeyWhitelist = make([]string, len(lines))
for i, line := range lines {
    publicKeyWhitelist[i] = strings.TrimRight(line, "\r\n")
 }

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
	return true
	//for _, k := range publicKeyWhitelist {
	//	if k == publicKey {
	//		return true
	//	}
	//}
	//return false
}

func Log(s string) {

}
