package database

import (
	"fmt"
	"log"
	"os"

	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/constant"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //pq is a pure Go Postgres driver for the database/sql package
)

const postgreDriver = "postgres"

// NewPostgresDatabase is used to create new Postgres setup
func NewPostgresDatabase(cfg *config.DatabaseConfig) *sqlx.DB {
	param := "sslmode=verify-full&sslrootcert=ap-southeast-1-bundle.pem&timezone=UTC"

	if os.Getenv(constant.EnvKeyEnv) != "production" {
		param = "sslmode=disable"
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		param,
	)

	db, err := sqlx.Open(postgreDriver, connStr)
	if err != nil {
		log.Fatal("failed to open db connection:", err.Error())
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	if err = db.Ping(); err != nil {
		log.Fatal("failed to ping db connection:", err.Error())
	}

	return db
}
