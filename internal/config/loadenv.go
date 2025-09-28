package config

import (
	"fmt"
)

// ServerConfig holds HTTP server settings
type ServerConfig struct {
	Host string
	Port int
}

func (s ServerConfig) Addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

// DBConfig holds database connection info
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// ConnString returns a ready-to-use PostgreSQL DSN
func (db DBConfig) ConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}

// AWSConfig holds AWS credentials/settings
type AWSConfig struct {
	AccessKey    string
	SecretKey    string
	Region       string
	Bucket       string
	SessionToken string
}

// JWTConfig holds JWT authentication settings
type JWTConfig struct {
	Secret     string
	Expiration int // in minutes
}

// -------------------- Root config --------------------

type config struct {
	Server ServerConfig
	DB     DBConfig
	AWS    AWSConfig
	JWT    JWTConfig
}

// Loadenv reads all environment variables into a Config struct
func Loadenv() config {
	return config{
		Server: ServerConfig{
			Host: getEnvString("SERVER_HOST", "127.0.0.1"),
			Port: getEnvInt("SERVER_PORT", 3000),
		},
		DB: DBConfig{
			Host:     getEnvString("DB_HOST", "127.0.0.1"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnvString("DB_USER", "postgres"),
			Password: getEnvString("DB_PASSWORD", "admin"),
			Name:     getEnvString("DB_NAME", "socialdb"),
		},
		AWS: AWSConfig{
			AccessKey:    getEnvString("AWS_ACCESS_KEY_ID", ""),
			SecretKey:    getEnvString("AWS_SECRET_ACCESS_KEY", ""),
			SessionToken: getEnvString("AWS_SESSION_TOKEN", ""),
			Region:       getEnvString("AWS_REGION", "us-east-1"),
			Bucket:       getEnvString("AWS_BUCKET", "my-bucket"),
		},
		JWT: JWTConfig{
			Secret:     getEnvString("JWT_SECRET", "secert23&^**&YHIi&T&Ghbkaknasdsft"),
			Expiration: getEnvInt("JWT_EXPIRATION", 3600), // 1h default
		},
	}
}
