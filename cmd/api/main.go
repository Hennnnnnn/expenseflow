package main

import (
	"log"
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/app"
	"github.com/Hennnnnnn/expenseflow/internal/config"
	"github.com/Hennnnnnn/expenseflow/internal/database"
	"github.com/Hennnnnnn/expenseflow/internal/email"
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
	"github.com/Hennnnnnn/expenseflow/internal/logger"
	"github.com/Hennnnnnn/expenseflow/internal/service"
	httptransport "github.com/Hennnnnnn/expenseflow/internal/transport/http"
	"github.com/Hennnnnnn/expenseflow/internal/transport/http/handlers"
)

func main() {

	// ==========================
	// Configuration
	// ==========================

	cfg := config.Load()
	logg := logger.New()

	application := &app.App{
		Config: cfg,
		Logger: logg,
	}

	// ==========================
	// Database
	// ==========================

	db, err := database.NewSQLite()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal(err)
	}

	application.DB = db

	// ==========================
	// Email
	// ==========================

	imapClient := imap.New(cfg)

	if err := imapClient.Connect(); err != nil {
		log.Fatal(err)
	}
	defer imapClient.Close()

	emailService := email.New(imapClient)
	emailHandler := handlers.NewEmailHandler(emailService)

	messages, err := emailService.ReadLatest(20)
	if err != nil {
		log.Fatal(err)
	}

	if len(messages) == 0 {
		log.Println("No BCA transaction emails found.")
	} else {
		data, err := emailService.Parse(&messages[0])
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("%+v\n", data)
		}

		body, err := emailService.ReadBody(messages[0].SeqNum)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("==========================")
		log.Println(body.TextBody)
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("========== Latest Emails ==========")

	for _, mail := range messages {
		log.Printf(
			"\nSubject : %s\nFrom    : %s\nDate    : %s\n",
			mail.Subject,
			mail.From,
			mail.Date,
		)
	}

	// ==========================
	// HTTP
	// ==========================

	transactionService := service.NewTransactionService(db)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	router := httptransport.NewRouter(
		logg,
		transactionHandler,
		emailHandler,
	)

	log.Println("🚀 ExpenseFlow API starting on :" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}
