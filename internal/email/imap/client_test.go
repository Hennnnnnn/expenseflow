package imap

import (
	"testing"

	"github.com/Hennnnnnn/expenseflow/internal/config"
)

func TestNewClient(t *testing.T) {
	cfg := &config.Config {
		IMAPHost: "imap.gmail.com",
		IMAPPort: "993",
		IMAPUsername: "test@gmail.com",
		IMAPPassword: "password",
	}

	client := New(cfg)

	if client.host != cfg.IMAPHost {
		t.Errorf("Expected host %s, got %s", cfg.IMAPHost, client.host)
	}

	if client.port != cfg.IMAPPort {
		t.Errorf("Expected port %s, got %s", cfg.IMAPPort, client.port)
	}
}