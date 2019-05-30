package smtp

import "github.com/vgxbj/mailinone/internal/config"

// Client ... Wrapper for smtp client.
type Client struct {
	account *config.Account
}

// NewClient ... Generate new smtp client.
func NewClient(acc *config.Account) (*Client, error) {
	return &Client{acc}, nil
}
