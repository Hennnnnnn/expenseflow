package imap

import "github.com/emersion/go-imap/v2"

func (c *Client) SelectInbox() (*imap.SelectData, error) {

	return c.client.
		Select("INBOX", nil).
		Wait()
}
