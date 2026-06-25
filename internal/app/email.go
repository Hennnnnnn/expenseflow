package app

import (
	"github.com/Hennnnnnn/expenseflow/internal/email"
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
)

func (a *App) BootstrapEmail() error {

	client := imap.New(a.Config)

	if err := client.Connect(); err != nil {
		return err
	}

	a.EmailService = email.New(client)

	return nil
}
