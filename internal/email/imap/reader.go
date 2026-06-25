package imap

import (
	goimap "github.com/emersion/go-imap/v2"
)

func (c *Client) ReadLatest(limit int) ([]Message, error) {

	selected, err := c.client.Select("INBOX", nil).Wait()
	if err != nil {
		return nil, err
	}

	total := selected.NumMessages

	if total == 0 {
		return []Message{}, nil
	}

	start := total - uint32(limit) + 1

	if total < uint32(limit) {
		start = 1
	}

	seqSet := goimap.SeqSetNum(start, total)

	fetchOptions := &goimap.FetchOptions{
		Envelope: true,
	}

	fetched, err := c.client.Fetch(seqSet, fetchOptions).Collect()

	if err != nil {
		return nil, err
	}

	var result []Message

	for _, msg := range fetched {

		var from string

		if len(msg.Envelope.From) > 0 {
			addr := msg.Envelope.From[0]

			from = addr.Mailbox + "@" + addr.Host
		}

		result = append(result, Message{
			SeqNum:  uint32(msg.SeqNum),
			Subject: msg.Envelope.Subject,
			From:    from,
			Date:    msg.Envelope.Date,
		})
	}

	return result, nil
}
