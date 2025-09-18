package config

import (
	"log"

	"github.com/joho/godotenv"
)

// AppConfig holds application-level settings
type AppConfig struct {
	Name     string
	Env      string
	Host     string
	Port     string
	LogLevel string
}

// DBConfig holds database connection settings
type PSQLConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// KafkaConfig holds Kafka settings
type KafkaConfig struct {
	Brokers             []string
	User                string
	Password            string
	ClientID            string
	TopicActiveAccount  string
	TopicForgotPassword string
}

// Redis config holds connect to redis
type RedisConfig struct {
	Host    string
	Port    int
	User    string
	Passwd  string
	SSLMode string
}

// Config is the root config struct
type Config struct {
	App   AppConfig
	PSQL  PSQLConfig
	Redis RedisConfig
	Kafka KafkaConfig
}

// Load loads environment variables and returns Config
func Load() *Config {
	// load .env file (ignore if missing, can use env vars directly)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables directly")
	}

	return &Config{
		App: AppConfig{
			Name:     getEnv("APP_NAME", "auth-service"),
			Env:      getEnv("APP_ENV", "development"),
			Host:     getEnv("APP_HOST", "127.0.0.1"),
			Port:     getEnv("APP_PORT", "8080"),
			LogLevel: getEnv("LOG_LEVEL", "info"),
		},
		PSQL: PSQLConfig{
			Host:     getEnv("PSQL_HOST", "localhost"),
			Port:     getEnv("PSQL_PORT", "5432"),
			User:     getEnv("PSQL_USER", "postgres"),
			Password: getEnv("PSQL_PASSWORD", ""),
			Name:     getEnv("PSQL_NAME", "auth_db"),
			SSLMode:  getEnv("PSQL_SSLMODE", "disable"),
		},
		Redis: RedisConfig{
			Host:    getEnv("REDIS_HOST", "localhost"),
			Port:    getEnvAsInt("REDIS_PORT", 6379),
			User:    getEnv("REDIS_USER", ""),
			Passwd:  getEnv("REDIS_PASSWD", ""),
			SSLMode: getEnv("REDIS_SSL_MODE", ""),
		},
		Kafka: KafkaConfig{
			Brokers:             splitEnv(getEnv("KAFKA_BROKERS", "localhost:9092")),
			User:                getEnv("KAFKA_USER", ""),
			Password:            getEnv("KAFKA_PASSWORD", ""),
			ClientID:            getEnv("KAFKA_CLIENT_ID", "auth-service-client"),
			TopicActiveAccount:  getEnv("KAFKA_TOPIC_USER_EVENTS", "active-users"),
			TopicForgotPassword: getEnv("TOPIC_FORGOTPASSWORD", "forgot-passwd"),
		},
	}
}
