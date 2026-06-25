package imap

import (
	"fmt"

	goimap "github.com/emersion/go-imap/v2"
)

func (c *Client) ReadLatest(limit int) ([]Message, error) {

	selected, err := c.client.Select("INBOX", nil).Wait()
	if err != nil {
		return nil, err
	}

	total := selected.NumMessages

	fmt.Printf("Total messages in inbox: %d\n", total)

	if total == 0 {
		return []Message{}, nil
	}

	var start uint32 = 1

	if total > uint32(limit) {
		start = total - uint32(limit) + 1
	}

	seqSet := goimap.SeqSet{
		goimap.SeqRange{
			Start: start,
			Stop:  total,
		},
	}

	fetchOptions := &goimap.FetchOptions{
		Envelope: true,
	}

	fmt.Println("Start :", start)

	fetched, err := c.client.Fetch(seqSet, fetchOptions).Collect()
	fmt.Printf("Fetched messages: %d\n", len(fetched))
	if err != nil {
		return nil, err
	}
	fmt.Println("End   :", total)

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
