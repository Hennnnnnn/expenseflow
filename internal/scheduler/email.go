package scheduler

import (
	"log"
	"time"

	"github.com/Hennnnnnn/expenseflow/internal/email"
)

type EmailScheduler struct {
	sync *email.SyncService

	interval time.Duration
}

func NewEmailScheduler(
	sync *email.SyncService,
	interval time.Duration,
) *EmailScheduler {

	return &EmailScheduler{
		sync:     sync,
		interval: interval,
	}
}

func (s *EmailScheduler) Start() {

	ticker := time.NewTicker(s.interval)

	go func() {

		for range ticker.C {

			log.Println("Running scheduled email sync...")

			result, err := s.sync.SyncLatest(100)

			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf(
				"Saved=%d Duplicate=%d Failed=%d Skipped=%d",
				result.Saved,
				result.Duplicate,
				result.Failed,
				result.Skipped,
			)

		}

	}()

}
