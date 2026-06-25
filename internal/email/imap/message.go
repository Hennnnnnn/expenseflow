package imap

import "time"

type Message struct {
	SeqNum uint32

	Subject string
	From    string
	Date    time.Time

	Header string

	TextBody string
	HTMLBody string
}
