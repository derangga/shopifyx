package config

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/derangga/shopifyx/internal/constant"
	"github.com/derangga/shopifyx/internal/pkg/helper"
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
	Auth        AuthConfig
}

func MustGet() *Config {
	once.Do(func() {
		cfg = &Config{}

		env := os.Getenv(constant.EnvKeyEnv)
		if env != "production" {
			godotenv.Load(".env")
			if err := cleanenv.ReadEnv(cfg); err != nil {
				log.Fatal("failed to load env:", err.Error())
			}

			return
		}

		cfg.Application = ApplicationConfig{
			Env:     env,
			Port:    helper.GetEnvWithDefault(constant.EnvKeyServicePort, "8000").ToString(),
			Service: helper.GetEnvWithDefault(constant.EnvKeyServiceName, "shopifyx").ToString(),
			Timeout: helper.GetEnvWithDefault(constant.EnvKeyServiceTimeout, "60s").ToDuration(),
		}

		cfg.Auth = AuthConfig{
			JWTSecret:        helper.GetEnvWithDefault(constant.EnvKeyJwtSecret, "").ToString(),
			JWTValidDuration: helper.GetEnvWithDefault(constant.EnvKeyJwtValidDuration, "2m").ToDuration(),
			BcryptSalt:       helper.GetEnvWithDefault(constant.EnvKeyBcryptSalt, "8").ToInt(),
		}

		cfg.Bucket = BucketConfig{
			ID:         helper.GetEnvWithDefault(constant.EnvKeyS3ID, "").ToString(),
			Secret:     helper.GetEnvWithDefault(constant.EnvKeyS3Secret, "").ToString(),
			BucketName: helper.GetEnvWithDefault(constant.EnvKeyS3BucketName, "").ToString(),
			Region:     helper.GetEnvWithDefault(constant.EnvKeyS3Region, "ap-southeast-3").ToString(),
			BaseURL:    helper.GetEnvWithDefault(constant.EnvKeyS3BaseURL, "").ToString(),
		}

		cfg.Database = DatabaseConfig{
			Name:         helper.GetEnvWithDefault(constant.EnvKeyDBName, "").ToString(),
			Port:         helper.GetEnvWithDefault(constant.EnvKeyDBPort, "").ToString(),
			Host:         helper.GetEnvWithDefault(constant.EnvKeyDBHost, "").ToString(),
			Username:     helper.GetEnvWithDefault(constant.EnvKeyDBUsername, "").ToString(),
			Password:     helper.GetEnvWithDefault(constant.EnvKeyDBPassword, "").ToString(),
			MaxOpenConns: helper.GetEnvWithDefault(constant.EnvKeyDBMaxOpenConns, "25").ToInt(),
			MaxIdleConns: helper.GetEnvWithDefault(constant.EnvKeyDBMaxidleConns, "25").ToInt(),
			MaxLifetime:  helper.GetEnvWithDefault(constant.EnvKeyDBMaxLifetime, "15m").ToDuration(),
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
	Region     string `env:"S3_REGION"      env-default:"ap-southeast-3"`
	BaseURL    string `env:"S3_BASE_URL"`
}

func (c *BucketConfig) ConstructURL() string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com", c.BucketName, c.Region)
}
