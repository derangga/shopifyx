package constant

const (
	EnvKeyEnv            = "ENV"
	EnvKeyServicePort    = "SERVICE_PORT"
	EnvKeyServiceName    = "SERVICE_NAME"
	EnvKeyServiceTimeout = "SERVICE_TIMEOUT"

	EnvKeyJwtSecret        = "JWT_SECRET"
	EnvKeyJwtValidDuration = "JWT_VALID_DURATION"
	EnvKeyBcryptSalt       = "BCRYPT_SALT"

	EnvKeyS3ID         = "S3_ID"
	EnvKeyS3Secret     = "S3_SECRET_KEY"
	EnvKeyS3BucketName = "S3_BUCKET_NAME"
	EnvKeyS3Region     = "S3_REGION"
	EnvKeyS3BaseURL    = "S3_BASE_URL"

	EnvKeyDBName         = "DB_NAME"
	EnvKeyDBPort         = "DB_PORT"
	EnvKeyDBHost         = "DB_HOST"
	EnvKeyDBUsername     = "DB_USERNAME"
	EnvKeyDBPassword     = "DB_PASSWORD"
	EnvKeyDBMaxOpenConns = "DB_MAX_OPEN_CONNS"
	EnvKeyDBMaxidleConns = "DB_MAX_IDLE_CONNS"
	EnvKeyDBMaxLifetime  = "DB_MAX_LIFETIME"
)
