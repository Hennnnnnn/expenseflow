package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Hennnnnnn/expenseflow/internal/ai"
	"github.com/Hennnnnnn/expenseflow/internal/ai/openai"
	"github.com/Hennnnnnn/expenseflow/internal/ai/pipeline"
	"github.com/Hennnnnnn/expenseflow/internal/app"
	"github.com/Hennnnnnn/expenseflow/internal/config"
	"github.com/Hennnnnnn/expenseflow/internal/database"
	"github.com/Hennnnnnn/expenseflow/internal/email"
	"github.com/Hennnnnnn/expenseflow/internal/email/imap"
	"github.com/Hennnnnnn/expenseflow/internal/logger"
	"github.com/Hennnnnnn/expenseflow/internal/scheduler"
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

	// ==========================
	// Services
	// ==========================

	transactionService := service.NewTransactionService(db)

	ruleEngine := ai.New()

	log.Println(cfg.OpenAIAPIKey)

	openAIProvider := openai.New(
		cfg.OpenAIAPIKey,
	)

	pipelineService := pipeline.New(
		ruleEngine,
		openAIProvider,
	)

	syncService := email.NewSyncService(
		emailService,
		email.NewProcessor(),
		transactionService,
		pipelineService,
	)

	emailScheduler := scheduler.NewEmailScheduler(
		syncService,
		10*time.Second,
	)

	emailScheduler.Start()

	// ==========================
	// Handlers
	// ==========================

	transactionHandler := handlers.NewTransactionHandler(transactionService)
	emailHandler := handlers.NewEmailHandler(syncService)

	// ==========================
	// Router
	// ==========================

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
