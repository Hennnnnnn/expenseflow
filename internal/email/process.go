package email

import (
	"errors"

	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
	"github.com/Hennnnnnn/expenseflow/internal/email/parser/bca"
)

func (s *SyncService) processMail(
	mail imap.Message,
	result *SyncResult,
) {

	body, err := s.reader.ReadBody(mail.SeqNum)
	if err != nil {
		println("❌ Failed to read body:", err.Error())
		result.Failed++
		return
	}

	transaction, err := s.processor.ProcessBytes(
		[]byte(body.TextBody),
	)

	if err != nil {

		if errors.Is(err, bca.ErrNotTransaction) {
			println("⏭ Skip non transaction email")
			result.Skipped++
			return
		}

		println("❌ Parse failed:", err.Error())
		result.Failed++
		return
	}

	saved, err := s.saveTransaction(transaction)
	if err != nil {
		println("❌ Save failed:", err.Error())
		result.Failed++
		return
	}

	if saved {
		result.Saved++
	} else {
		result.Duplicate++
	}
}
