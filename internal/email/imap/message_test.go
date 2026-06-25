package imap

import "testing"

func TestMessage(t *testing.T) {
	msg := Message{
		Subject: "BCA Notification",
	}

	if msg.Subject != "BCA Notification" {
		t.Fatal("invalid subject")
	}
}
