package imap

import (
	"crypto/tls"

	"github.com/emersion/go-imap/client"
	"github.com/vgxbj/mailinone/internal/config"
)

// Dial ... Dial imap mail server according to mailbox configuration.
func Dial(mb *config.Mailbox) (*client.Client, error) {
	if mb.EnableSSL() {
		tlsConfig := &tls.Config{InsecureSkipVerify: true}
		return client.DialTLS(mb.IncomingMailServer().Addr(), tlsConfig)
	}

	return client.Dial(mb.IncomingMailServer().Addr())
}
