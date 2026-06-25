package main

import (
	"log"
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/app"
	"github.com/Hennnnnnn/expenseflow/internal/config"
	"github.com/Hennnnnnn/expenseflow/internal/logger"
	"github.com/Hennnnnnn/expenseflow/internal/database"
	httptransport "github.com/Hennnnnnn/expenseflow/internal/transport/http"
)

func main() {

	cfg := config.Load()

	logg := logger.New()

	application := &app.App{
		Config: cfg,
		Logger: logg,
	}

	db, err := database.NewSQLite()

	if err != nil {
		log.Fatal(err)
	}

	application.DB = db

	router := httptransport.NewRouter(logg)
	
	log.Println("🚀 ExpenseFlow API starting on :" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}