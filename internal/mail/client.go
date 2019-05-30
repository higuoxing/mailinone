package mail

import (
	"github.com/vgxbj/mailinone/internal/config"
	"github.com/vgxbj/mailinone/internal/mail/imap"
	"github.com/vgxbj/mailinone/internal/mail/smtp"
)

// Client ... Wrapper for imap, smtp client.
type Client struct {
	imapClient *imap.Client
	smtpClient *smtp.Client
}

// NewClient ... Generate new client.
func NewClient(acc *config.Account) (*Client, error) {
	imapClient, err := imap.NewClient(acc)
	if err != nil {
		return nil, err
	}

	smtpClient, err := smtp.NewClient(acc)
	if err != nil {
		return nil, err
	}

	return &Client{imapClient, smtpClient}, nil
}
