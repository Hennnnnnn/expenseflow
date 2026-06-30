package email

// func TestSyncFile(t *testing.T) {

// 	db, err := database.NewSQLite()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := database.AutoMigrate(db); err != nil {
// 		t.Fatal(err)
// 	}

// 	transactionService := service.NewTransactionService(db)

// 	processor := NewProcessor()

// 	sync := NewSyncService(
// 		nil,
// 		processor,
// 		transactionService,
// 		ai.New(),
// 	)

// 	err = sync.SyncFile(
// 		"../../testdata/bca/credit_card_transaction.eml",
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
