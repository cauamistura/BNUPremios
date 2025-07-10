package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config representa as configurações da aplicação
type Config struct {
	Database DatabaseConfig
	API      APIConfig
	JWT      JWTConfig
}

// DatabaseConfig representa as configurações do banco de dados
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// APIConfig representa as configurações da API
type APIConfig struct {
	Port string
	Mode string
}

// JWTConfig representa as configurações do JWT
type JWTConfig struct {
	Secret string
}

// Load carrega as configurações da aplicação
func Load() *Config {
	// Carregar arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Name:     getEnv("DB_NAME", "bnupremios"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		API: APIConfig{
			Port: getEnv("API_PORT", "8080"),
			Mode: getEnv("API_MODE", "debug"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "your-secret-key-here"),
		},
	}
}

// getEnv obtém uma variável de ambiente ou retorna um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}