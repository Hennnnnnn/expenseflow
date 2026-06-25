package main

import (
	"log"
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/app"
	"github.com/Hennnnnnn/expenseflow/internal/config"
	"github.com/Hennnnnnn/expenseflow/internal/database"
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
	"github.com/Hennnnnnn/expenseflow/internal/logger"
	"github.com/Hennnnnnn/expenseflow/internal/service"
	httptransport "github.com/Hennnnnnn/expenseflow/internal/transport/http"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/handlers"
)

func main() {

	cfg := config.Load()

	imapClient := imap.New(cfg)

	emailService := service.New(imapClient)

	_ = emailService

	if err := imapClient.Connect(); err != nil {
		log.Fatal(err)
	}

	defer imapClient.Close()

	mailbox, err := imapClient.SelectInbox()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Inbox has %d messages\n", mailbox.NumMessages)

	log.Println("✅ Connected to Gmail IMAP")

	logg := logger.New()

	application := &app.App{
		Config: cfg,
		Logger: logg,
	}

	db, err := database.NewSQLite()

	if err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}

	application.DB = db

	transactionService := service.NewTransactionService(db)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	_ = transactionService

	router := httptransport.NewRouter(logg, transactionHandler)

	log.Println("🚀 ExpenseFlow API starting on :" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
