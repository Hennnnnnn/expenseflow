package imap

import (
	"fmt"

	"github.com/Hennnnnnn/expenseflow/internal/config"
	"github.com/emersion/go-imap/v2/imapclient"
)

type Client struct {
	host     string
	port     string
	username string
	password string

	client *imapclient.Client
}

func New(cfg *config.Config) *Client {
	return &Client{
		host:     cfg.IMAPHost,
		port:     cfg.IMAPPort,
		username: cfg.IMAPUsername,
		password: cfg.IMAPPassword,
	}
}

func (c *Client) Connect() error {
	address := fmt.Sprintf("%s:%s", c.host, c.port)

	client, err := imapclient.DialTLS(
		address,
		&imapclient.Options{},
	)
	if err != nil {
		return err
	}

	if err := client.Login(c.username, c.password).Wait(); err != nil {
		return err
	}

	c.client = client

	return nil
}

func (c *Client) Close() error {
	if c.client == nil {
		return nil
	}

	return c.client.Logout().Wait()
}

func (c *Client) ReadLatest(limit int) ([]Message, error) {
	return nil, nil
}
