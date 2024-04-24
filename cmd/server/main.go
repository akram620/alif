package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/akram620/alif/internal/api"
	"github.com/akram620/alif/internal/config"
	"github.com/akram620/alif/internal/logger"
	"github.com/akram620/alif/internal/migrate"
	"github.com/akram620/alif/internal/repository"
	"github.com/akram620/alif/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

func main() {
	if err := config.LoadFromFile(".env"); err != nil {
		logger.Fatalf("config.LoadFromFile(): %v", err)
	}

	pool, err := connectToDatabase(config.Values.DatabaseURL)
	if err != nil {
		logger.Fatalf("connectToDatabase(): %v", err)
	}
	defer pool.Close()

	if err := migrate.ApplyMigrations("migrations"); err != nil {
		logger.Fatalf("migrate.ApplyMigrations(): %v", err)
	}

	chatRepository := repository.NewEventsRepository(pool)

	chatService := service.NewEventsService(chatRepository)

	server := api.NewServer(chatService)
	server.Run(config.Values.APIPort)
}

func connectToDatabase(url string) (*pgxpool.Pool, error) {
	var retries int
	var maxRetries = 5

	var pool *pgxpool.Pool
	var err error

	if len(url) == 0 {
		return nil, errors.New("missing DB_URL environment variable")
	}

	for {
		if retries >= maxRetries {
			return nil, fmt.Errorf("couldn't connect to the database after %d retries", retries)
		}

		pool, err = pgxpool.New(context.Background(), url)
		if err != nil {
			logger.Errorf("couldn't connect to the database: %v", err)
			time.Sleep(2 * time.Second)

			retries++
			continue
		}

		logger.Infof("successfully connected")
		return pool, nil
	}
}
