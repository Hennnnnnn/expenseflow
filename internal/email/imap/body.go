package imap

import (
	"bytes"

	goimap "github.com/emersion/go-imap/v2"
)

func (c *Client) ReadBody(seqNum uint32) (*Message, error) {

	seqSet := goimap.SeqSetNum(seqNum)

	bodySection := &goimap.FetchItemBodySection{}

	fetchOptions := &goimap.FetchOptions{
		Envelope: true,
		BodySection: []*goimap.FetchItemBodySection{
			bodySection,
		},
	}

	fetched, err := c.client.Fetch(seqSet, fetchOptions).Collect()

	if err != nil {
		return nil, err
	}

	if len(fetched) == 0 {
		return nil, nil
	}

	msg := fetched[0]

	body := msg.FindBodySection(bodySection)

	return &Message{
		SeqNum:   seqNum,
		Subject:  msg.Envelope.Subject,
		Date:     msg.Envelope.Date,
		TextBody: string(bytes.TrimSpace(body)),
	}, nil
}
