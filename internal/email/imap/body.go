package imap

import "fmt"

func (c *Client) ReadBody(uid uint32) (*Message, error) {
	return nil, fmt.Errorf("ReadBody not implemented")
}
