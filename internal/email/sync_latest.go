package email

func (s *Service) SyncLatest(limit int) error {

	emails, err := s.ReadLatest(limit)

	if err != nil {
		return err
	}

	for _, mail := range emails {

		println(mail.Subject)

	}

	return nil
}
