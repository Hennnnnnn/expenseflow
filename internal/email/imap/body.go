package imap

import (
	goimap "github.com/emersion/go-imap/v2"
)

func (c *Client) ReadBody(seqNum uint32) (*Message, error) {

	seqSet := goimap.SeqSetNum(seqNum)

	headerSection := &goimap.FetchItemBodySection{
		Specifier: goimap.PartSpecifierHeader,
	}

	fetchOptions := &goimap.FetchOptions{
		Envelope: true,
		BodySection: []*goimap.FetchItemBodySection{
			headerSection,
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

	header := msg.FindBodySection(headerSection)

	return &Message{
		UID:     seqNum,
		Subject: msg.Envelope.Subject,
		Date:    msg.Envelope.Date,
		Header:  string(header),
	}, nil
}
