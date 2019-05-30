package config

import (
	"fmt"
	"strconv"

	"github.com/BurntSushi/toml"
)

// Configs ... Struct of mailinone configuration.
type Configs struct {
	Accounts []Account `toml:"Account"` // account configuration
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
	var accs []Account

	accs = c.Accounts

	if len(accs) == 0 {
		return fmt.Errorf("At least one [[Account]] is required")
	}

	for _, acc := range accs {
		if err = acc.Verify(); err != nil {
			return err
		}
	}

	return nil
}

// Account ... Struct for account configuration.
type Account struct {
	Username                    string             `toml:"Username"`                    // username for login
	Password                    string             `toml:"Password"`                    // password for login
	EnableSSL                   bool               `toml:"EnableSSL"`                   // enable SSL when login email
	SkipCertificateVerification bool               `toml:"SkipCertificateVerification"` // skip ssl certificate verification
	IncomingMailServer          IncomingMailServer `toml:"IncomingMailServer"`          // incoming mail server
	OutgoingMailServer          OutgoingMailServer `toml:"OutgoingMailServer"`          // outgoing mail server
}

// Verify ... Check account configuration.
func (acc *Account) Verify() error {
	var err error

	if acc.Username == "" {
		return fmt.Errorf("Username is required in [[Account]]")
	}

	if acc.Password == "" {
		return fmt.Errorf("Password is required in [[Account]]")
	}

	if err = acc.IncomingMailServer.Verify(); err != nil {
		return err
	}

	if err = acc.OutgoingMailServer.Verify(); err != nil {
		return err
	}

	return nil
}

// IncomingMailServer ... Struct of incoming mail server configuration.
type IncomingMailServer struct {
	Hostname string `toml:"Hostname"` // incoming mail hostname
	Port     int    `toml:"Port"`     // incoming mail server port
}

// Addr ... Get incoming mail server address.
func (ims *IncomingMailServer) Addr() string {
	return ims.Hostname + ":" + strconv.Itoa(ims.Port)
}

// Verify ... Check incoming mail server configuration.
func (ims *IncomingMailServer) Verify() error {
	if ims.Hostname == "" {
		return fmt.Errorf("Hostname is required in [Account.IncomingMailServer]")
	}

	if ims.Port == 0 {
		return fmt.Errorf("Port is required in [Account.IncomingMailServer] and cannot be set to '0'")
	}

	return nil
}

// OutgoingMailServer ... Struct of outgoing mail server configuration.
type OutgoingMailServer struct {
	Hostname string `toml:"Hostname"` // outgoing mail hostname
	Port     int    `toml:"Port"`     // outgoing mail server port
}

// Addr ... Get outgoing mail server address.
func (oms *OutgoingMailServer) Addr() string {
	return oms.Hostname + ":" + strconv.Itoa(oms.Port)
}

// Verify ... Check outgoing mail server configuration.
func (oms *OutgoingMailServer) Verify() error {
	if oms.Hostname == "" {
		return fmt.Errorf("Hostname is required in [Account.OutgoingMailServer]")
	}

	if oms.Port == 0 {
		return fmt.Errorf("Port is required in [Account.OutgoingMailServer] and cannot be set to '0'")
	}

	return nil
}
