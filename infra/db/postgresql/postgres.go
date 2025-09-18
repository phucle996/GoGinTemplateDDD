package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"authen_service/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPostgres khởi tạo kết nối PostgreSQL với sqlx và retry
func NewPostgres(cfg *config.Config) *sqlx.DB {
	var db *sqlx.DB
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PSQL.Host, cfg.PSQL.Port, cfg.PSQL.User, cfg.PSQL.Password, cfg.PSQL.Name, cfg.PSQL.SSLMode,
	)

	maxRetries := 5
	retryInterval := 3 * time.Second

	for i := 1; i <= maxRetries; i++ {
		db, err = sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Printf("Attempt %d: failed to connect DB: %v", i, err)
		} else {
			// Ping với context timeout 3s
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()
			if pingErr := db.PingContext(ctx); pingErr != nil {
				log.Printf("Attempt %d: failed to ping DB: %v", i, pingErr)
				err = pingErr
			} else {
				log.Println("Connected to PostgreSQL successfully")
				return db
			}
		}

		if i < maxRetries {
			log.Printf("Retrying in %s...", retryInterval)
			time.Sleep(retryInterval)
		}
	}

	log.Fatalf("Could not connect to PostgreSQL after %d attempts: %v", maxRetries, err)
	return nil
}
