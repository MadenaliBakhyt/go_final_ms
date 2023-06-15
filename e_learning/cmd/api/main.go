package main

import (
	"e-learning/internal/data"
	db "e-learning/internal/db"
	"e-learning/internal/jsonlog"
	"e-learning/pkg/config"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type application struct {
	config *config.Config
	logger *jsonlog.Logger
	models data.Models
	wg     sync.WaitGroup
}

func main() {
	var cfg = config.GetConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	db, err := db.OpenDB()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()
	logger.PrintInfo("database connection pool established", nil)

	app := &application{
		config: &cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
