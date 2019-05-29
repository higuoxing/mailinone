package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Configs ... Struct of mailinone configuration.
type Configs struct {
	Mailboxes []mailbox `toml:"mailbox"` // mails
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

	if len(c.Mailboxes) == 0 {
		return fmt.Errorf("At least one [[mailbox]] is required")
	}

	for _, mb := range c.Mailboxes {
		if err = mb.Verify(); err != nil {
			return err
		}
	}

	return nil
}

type mailbox struct {
	Username            string              `toml:"username"`  // username for login
	Password            string              `toml:"password"`  // password for login
	EnableSSL           bool                `toml:"enableSSL"` // enable SSL when login email
	IncommingMailServer incommingMailServer `toml:"incomming"` // incomming mail server
	OutgoingMailServer  outgoingMailServer  `toml:"outgoing"`  // outgoing mail server
}

// Verify ... Check mailbox configuration.
func (mb *mailbox) Verify() error {
	var err error

	if mb.Username == "" {
		return fmt.Errorf("Username is required in [[mailbox]] field")
	}

	if mb.Password == "" {
		return fmt.Errorf("Password is required in [[mailbox]] field")
	}

	if err = mb.IncommingMailServer.Verify(); err != nil {
		return err
	}

	if err = mb.OutgoingMailServer.Verify(); err != nil {
		return err
	}

	return nil
}

type incommingMailServer struct {
	Hostname string `toml:"hostname"` // incomming mail hostname
	Port     int    `toml:"port"`     // incomming mail server port
}

// Verify ... Check incomming mail server configuration.
func (ims *incommingMailServer) Verify() error {
	if ims.Hostname == "" {
		return fmt.Errorf("Hostname is required in [mailbox.incomming] field")
	}

	if ims.Port == 0 {
		ims.Port = 993 // set it to imap's default port
	}

	return nil
}

type outgoingMailServer struct {
	Hostname string `toml:"hostname"` // outgoing mail hostname
	Port     int    `toml:"port"`     // outgoing mail server port
}

// Verify ... Check outgoing mail server configuration.
func (oms *outgoingMailServer) Verify() error {
	if oms.Hostname == "" {
		return fmt.Errorf("Hostname is required in [mailbox.outgong] field")
	}

	if oms.Port == 0 {
		oms.Port = 465 // set it to smtp's default port
	}

	return nil
}
