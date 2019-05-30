package imap

import (
	"crypto/tls"

	"github.com/emersion/go-imap"

	"github.com/emersion/go-imap/client"
	"github.com/vgxbj/mailinone/internal/config"
)

// Client ... Wrapper for imap client.
type Client struct {
	account *config.Account
	client  *client.Client
}

// MailboxInfo ... Alias for imap.MailboxInfo
type MailboxInfo = imap.MailboxInfo

// NewClient ... Generate new imap client.
func NewClient(acc *config.Account) (*Client, error) {
	var c *client.Client
	var err error

	if acc.EnableSSL {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: acc.SkipCertificateVerification,
		}

		c, err = client.DialTLS(acc.IncomingMailServer.Addr(), tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		c, err = client.Dial(acc.IncomingMailServer.Addr())
		if err != nil {
			return nil, err
		}
	}

	return &Client{acc, c}, nil
}

// Login ... Login imap mail server.
func (c *Client) Login() error {
	return c.client.Login(c.account.Username, c.account.Password)
}

// Logout ... Logout imap mail server.
func (c *Client) Logout() error {
	return c.client.Logout()
}

// GetMailboxes ... Get mailboxes of current account.
func (c *Client) GetMailboxes() ([]*MailboxInfo, error) {
	var mbs []*MailboxInfo
	var mbch = make(chan *MailboxInfo, 5)
	var done = make(chan error, 1)

	go func() {
		done <- c.client.List("", "*", mbch)
	}()

	for mb := range mbch {
		mbs = append(mbs, mb)
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return mbs, nil
}
