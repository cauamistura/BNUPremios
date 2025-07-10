package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cauamistura/BNUPremios/internal/models"
	"github.com/google/uuid"
)

// UserRepository implementa as operações de banco de dados para usuários
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository cria uma nova instância do repositório de usuários
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create cria um novo usuário
func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (id, name, email, password, role, active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if user.Role == "" {
		user.Role = "user"
	}

	_, err := r.db.Exec(query,
		user.ID, user.Name, user.Email, user.Password, user.Role, user.Active,
		user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return nil
}

// GetByID busca um usuário pelo ID
func (r *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT id, name, email, password, role, active, created_at, updated_at
		FROM users WHERE id = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Role,
		&user.Active, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

// GetByEmail busca um usuário pelo email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, email, password, role, active, created_at, updated_at
		FROM users WHERE email = $1
	`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Role,
		&user.Active, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}

	return user, nil
}

// List busca todos os usuários com paginação
func (r *UserRepository) List(page, limit int) ([]models.User, int, error) {
	// Contar total de registros
	countQuery := `SELECT COUNT(*) FROM users`
	var total int
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao contar usuários: %w", err)
	}

	// Calcular offset
	offset := (page - 1) * limit

	// Buscar usuários
	query := `
		SELECT id, name, email, password, role, active, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao listar usuários: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Password, &user.Role,
			&user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, 0, fmt.Errorf("erro ao escanear usuário: %w", err)
		}
		users = append(users, user)
	}

	return users, total, nil
}

// Update atualiza um usuário
func (r *UserRepository) Update(id uuid.UUID, updates map[string]interface{}) error {
	// Construir query dinamicamente
	query := `UPDATE users SET `
	args := []interface{}{}
	argIndex := 1

	for field, value := range updates {
		if argIndex > 1 {
			query += ", "
		}
		query += fmt.Sprintf("%s = $%d", field, argIndex)
		args = append(args, value)
		argIndex++
	}

	query += fmt.Sprintf(", updated_at = $%d WHERE id = $%d", argIndex, argIndex+1)
	args = append(args, time.Now(), id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}

	return nil
}

// Delete remove um usuário
func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("usuário não encontrado")
	}

	return nil
}

// EmailExists verifica se um email já existe
func (r *UserRepository) EmailExists(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	
	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar existência do email: %w", err)
	}

	return exists, nil
} 