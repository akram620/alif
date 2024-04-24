package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/akram620/alif/internal/config"
	"github.com/akram620/alif/internal/handler"
	"github.com/akram620/alif/internal/logger"
	"github.com/akram620/alif/internal/migrate"
	"github.com/akram620/alif/internal/repository"
	"github.com/akram620/alif/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	repeatInterval = time.Minute
)

func main() {
	// загружаем переменные из файла или подтягиваем из процесса (на сервере) , сохраняем в структуре
	if err := config.LoadFromFile(".env"); err != nil {
		logger.Fatalf("config.LoadFromFile(): %v", err)
	}

	pool, err := connectToDatabase(config.Values.DatabaseURL)
	if err != nil {
		logger.Fatalf("connectToDatabase(): %v", err)
	}

	if err := migrate.ApplyMigrations("schema"); err != nil {
		logger.Fatalf("migrate.ApplyMigrations(): %v", err)
	}

	// инициализируем зависимости
	eventsRepository := repository.NewEventsRepository(pool)
	eventsService := service.NewEventsService(eventsRepository)

	ctx, cancel := context.WithCancel(context.Background())
	workerService := service.NewWorkerService(eventsRepository)
	go workerService.RunJobs(ctx, repeatInterval)

	server := handler.NewHandler(eventsService)
	go server.Run(config.Values.APIPort)

	awaitQuitSignal(cancel, pool)
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

func awaitQuitSignal(cancel context.CancelFunc, pool *pgxpool.Pool) {
	logger.Infof("Server started. Working until a quit signal is received...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	cancel()
	pool.Close()
	logger.Infof("Stopping server...")
}
