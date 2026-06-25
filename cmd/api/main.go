package main

import (
	"log"
	"net/http"

	"github.com/Hennnnnnn/expenseflow/internal/app"
	"github.com/Hennnnnnn/expenseflow/internal/config"
	"github.com/Hennnnnnn/expenseflow/internal/logger"
	httptransport "github.com/Hennnnnnn/expenseflow/internal/transport/http"
)

func main() {

	cfg := config.Load()

	logg := logger.New()

	application := &app.App{
		Config: cfg,
		Logger: logg,
	}

	_ = application

	router := httptransport.NewRouter(logg)
	
	log.Println("🚀 ExpenseFlow API starting on :" + cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal(err)
	}
}