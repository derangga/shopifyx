package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

var (
	once sync.Once
	cfg  *Config
)

// Config holds all configuration
type Config struct {
	Application ApplicationConfig
	Database    DatabaseConfig
	Bucket      BucketConfig
	Monitoring  MonitoringConfig
	Auth        AuthConfig
}

func MustGet() *Config {
	once.Do(func() {
		cfg = &Config{}

		godotenv.Load(".env")
		if err := cleanenv.ReadEnv(cfg); err != nil {
			log.Fatal("failed to load env:", err.Error())
		}
	})

	return cfg
}

// ApplicationConfig holds the configuration for the app instance
type ApplicationConfig struct {
	Env     string        `env:"ENV"             env-default:"local"`
	Port    string        `env:"SERVICE_PORT"    env-default:"8000"`
	Service string        `env:"SERVICE_NAME"    env-default:"shopifyx"`
	Timeout time.Duration `env:"SERVICE_TIMEOUT" env-default:"60s"`
}

// DatabaseConfig holds the configuration for the database instance
type DatabaseConfig struct {
	Name         string        `env:"DB_NAME"`
	Port         string        `env:"DB_PORT"`
	Host         string        `env:"DB_HOST"`
	Username     string        `env:"DB_USERNAME"`
	Password     string        `env:"DB_PASSWORD"`
	MaxOpenConns int           `env:"DB_MAX_OPEN_CONNS" env-default:"25"`
	MaxIdleConns int           `env:"DB_MAX_IDLE_CONNS" env-default:"25"`
	MaxLifetime  time.Duration `env:"DB_MAX_LIFETIME"   env-default:"15m"`
}

// AuthConfig holds the configuration for auth
type AuthConfig struct {
	JWTSecret        string        `env:"JWT_SECRET"`
	JWTValidDuration time.Duration `env:"JWT_VALID_DURATION" env-default:"2m"`
	BcryptSalt       int           `env:"BCRYPT_SALT"        env-default:"8"`
}

// BucketConfig holds the configuration for bucket
type BucketConfig struct {
	ID         string `env:"S3_ID"`
	Secret     string `env:"S3_SECRET_KEY"`
	BucketName string `env:"S3_BUCKET_NAME" env-default:"sprint-bucket-public-read"`
	Region     string `env:"S3_REGION"      env-default:"ap-southeast-1"`
	BaseURL    string `env:"S3_BASE_URL"`
}

func (c *BucketConfig) ConstructURL() string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com", c.BucketName, c.Region)
}

// MonitoringConfig holds the configuration for monitoring
type MonitoringConfig struct {
	PrometheusAddrs string `env:"PROMETHEUS_ADDRESS"`
}
