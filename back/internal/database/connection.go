package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cauamistura/BNUPremios/internal/config"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Connect estabelece conexão com o banco de dados PostgreSQL
func Connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com banco de dados: %w", err)
	}

	// Testar conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao conectar com banco de dados: %w", err)
	}

	log.Println("Conexão com banco de dados estabelecida com sucesso")
	return db, nil
}

// RunMigrations executa as migrações do banco de dados
func RunMigrations(cfg config.DatabaseConfig) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return fmt.Errorf("erro ao criar migrator: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("erro ao executar migrações: %w", err)
	}

	log.Println("Migrações executadas com sucesso")
	return nil
} 