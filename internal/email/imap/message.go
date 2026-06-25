package imap

import "time"

type Message struct {
	UID     uint32
	Subject string
	From    string
	Date    time.Time
}
