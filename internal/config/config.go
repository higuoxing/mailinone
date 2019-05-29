package config

import (
	"fmt"
	"strconv"

	"github.com/BurntSushi/toml"
)

// Configs ... Struct of mailinone configuration.
type Configs struct {
	mailboxes []Mailbox `toml:"Mailbox"` // mails
}

// Mailboxes ... Get mailboxes' configuration.
func (c *Configs) Mailboxes() []Mailbox {
	return c.mailboxes
}

// ReadConfigFromFile ... Read configuration from given file.
func ReadConfigFromFile(p string) (*Configs, error) {
	var err error
	var configs Configs

	if _, err = toml.DecodeFile(p, &configs); err != nil {
		return nil, err
	}

	return &configs, nil
}

// Verify ... Check configuration.
func (c *Configs) Verify() error {
	var err error
	var mbs []Mailbox

	mbs = c.Mailboxes()

	if len(mbs) == 0 {
		return fmt.Errorf("At least one [[Mailbox]] is required")
	}

	for _, mb := range mbs {
		if err = mb.Verify(); err != nil {
			return err
		}
	}

	return nil
}

// Mailbox ... Struct for mailbox configuration.
type Mailbox struct {
	username                    string             `toml:"Username"`                    // username for login
	password                    string             `toml:"Password"`                    // password for login
	enableSSL                   bool               `toml:"EnableSSL"`                   // enable SSL when login email
	skipCertificateVerification bool               `toml:"SkipCertificateVerification"` // skip ssl certificate verification
	incomingServer              IncomingMailServer `toml:"IncomingMailServer"`          // incoming mail server
	outgoingServer              OutgoingMailServer `toml:"OutgoingMailServer"`          // outgoing mail server
}

// Username ... Get username of mailbox.
func (mb *Mailbox) Username() string {
	return mb.username
}

// Password ... Get password of mailbox.
func (mb *Mailbox) Password() string {
	return mb.password
}

// EnableSSL ... Check if mailbox has SSL enabled.
func (mb *Mailbox) EnableSSL() bool {
	return mb.enableSSL
}

// IsSkipCertVerification ... Check if mailbox has certificate verification disabled.
func (mb *Mailbox) IsSkipCertVerification() bool {
	return mb.skipCertificateVerification
}

// IncomingMailServer ... Get incoming mail server configuration.
func (mb *Mailbox) IncomingMailServer() *IncomingMailServer {
	return &mb.incomingServer
}

// OutgoingMailServer ... Get outgoing mail server configuration.
func (mb *Mailbox) OutgoingMailServer() *OutgoingMailServer {
	return &mb.outgoingServer
}

// Verify ... Check mailbox configuration.
func (mb *Mailbox) Verify() error {
	var err error

	if mb.Username() == "" {
		return fmt.Errorf("Username is required in [[Mailbox]]")
	}

	if mb.Password() == "" {
		return fmt.Errorf("Password is required in [[Mailbox]]")
	}

	if err = mb.IncomingMailServer().Verify(); err != nil {
		return err
	}

	if err = mb.OutgoingMailServer().Verify(); err != nil {
		return err
	}

	return nil
}

// IncomingMailServer ... Struct of incoming mail server configuration.
type IncomingMailServer struct {
	hostname string `toml:"Hostname"` // incoming mail hostname
	port     int    `toml:"Port"`     // incoming mail server port
}

// Hostname ... Get hostname of incoming mail server.
func (ims *IncomingMailServer) Hostname() string {
	return ims.hostname
}

// Port ... Get port of incoming mail server.
func (ims *IncomingMailServer) Port() int {
	return ims.port
}

// Addr ... Get incoming mail server address.
func (ims *IncomingMailServer) Addr() string {
	return ims.Hostname() + ":" + strconv.Itoa(ims.Port())
}

// Verify ... Check incoming mail server configuration.
func (ims *IncomingMailServer) Verify() error {
	if ims.Hostname() == "" {
		return fmt.Errorf("Hostname is required in [Mailbox.IncomingMailServer]")
	}

	if ims.Port() == 0 {
		return fmt.Errorf("Port is required in [Mailbox.IncomingMailServer] and cannot be set to '0'")
	}

	return nil
}

// OutgoingMailServer ... Struct of outgoing mail server configuration.
type OutgoingMailServer struct {
	hostname string `toml:"Hostname"` // outgoing mail hostname
	port     int    `toml:"Port"`     // outgoing mail server port
}

// Hostname ... Get hostname of outgoing mail server.
func (oms *OutgoingMailServer) Hostname() string {
	return oms.hostname
}

// Port ... Get port of outgoing mail server.
func (oms *OutgoingMailServer) Port() int {
	return oms.port
}

// Addr ... Get outgoing mail server address.
func (oms *OutgoingMailServer) Addr() string {
	return oms.Hostname() + ":" + strconv.Itoa(oms.Port())
}

// Verify ... Check outgoing mail server configuration.
func (oms *OutgoingMailServer) Verify() error {
	if oms.Hostname() == "" {
		return fmt.Errorf("Hostname is required in [Mailbox.OutgoingMailServer]")
	}

	if oms.Port() == 0 {
		return fmt.Errorf("Port is required in [Mailbox.OutgoingMailServer] and cannot be set to '0'")
	}

	return nil
}
