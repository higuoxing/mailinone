package imap

import (
	"crypto/tls"
	"fmt"

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

// MailboxStatus ... Alias for imap.MailboxStatus
type MailboxStatus = imap.MailboxStatus

// Message ... Alias for imap.Message
type Message = imap.Message

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
	var mbch = make(chan *MailboxInfo)
	var done = make(chan error, 1)

	go func() {
		done <- c.client.List("", "*", mbch)
	}()

	// Append mailboxes to slice.
	for mb := range mbch {
		mbs = append(mbs, mb)
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return mbs, nil
}

// GetMailbox ... Get mailbox of given name.
func (c *Client) GetMailbox(name string) (*MailboxStatus, error) {
	mb, err := c.client.Select(name, true /* Readonly mode */)
	if err != nil {
		return nil, err
	}

	defer c.client.Close()

	return mb, nil
}

// FetchMailsOf ... Fetch mails of given mailbox.
func (c *Client) FetchMailsOf(name string, from, delta uint32) ([]*Message, error) {
	mb, err := c.client.Select(name, true /* Readonly mode */)
	if err != nil {
		return nil, err
	}

	defer c.client.Close()

	if from > mb.Messages {
		return nil, fmt.Errorf("Mailbox %s only have %d messages", name, int(mb.Messages))
	}

	if from+delta > mb.Messages {
		delta = mb.Messages - from
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(from, from+delta)

	msgch := make(chan *Message, 10)
	done := make(chan error)

	go func() {
		done <- c.client.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, msgch)
	}()

	var msgs []*Message

	for m := range msgch {
		msgs = append(msgs, m)
	}

	if err := <-done; err != nil {
		return nil, err
	}

	return msgs, nil
}
