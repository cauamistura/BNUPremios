package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cauamistura/BNUPremios/internal/models"
	"github.com/google/uuid"
)

type RewardRepository struct {
	db *sql.DB
}

func NewRewardRepository(db *sql.DB) *RewardRepository {
	return &RewardRepository{db: db}
}

// Create cria um novo prêmio
func (r *RewardRepository) Create(reward *models.Reward, price float64, minQuota int, images []string) error {
	// Iniciar transação
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Inserir prêmio básico
	rewardQuery := `
		INSERT INTO rewards (id, name, description, image, draw_date, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = tx.Exec(rewardQuery,
		reward.ID,
		reward.Name,
		reward.Description,
		reward.Image,
		reward.DrawDate,
		reward.Completed,
		reward.CreatedAt,
		reward.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// Inserir detalhes do prêmio
	detailsQuery := `
		INSERT INTO reward_details (reward_id, price, min_quota, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = tx.Exec(detailsQuery,
		reward.ID,
		price,
		minQuota,
		reward.CreatedAt,
		reward.UpdatedAt,
	)
	if err != nil {
		return err
	}

	// Inserir imagens adicionais
	if len(images) > 0 {
		imagesQuery := `
			INSERT INTO reward_images (reward_id, image_url, created_at)
			VALUES ($1, $2, $3)
		`

		for _, image := range images {
			_, err = tx.Exec(imagesQuery, reward.ID, image, reward.CreatedAt)
			if err != nil {
				return err
			}
		}
	}

	// Commit da transação
	return tx.Commit()
}

// GetByID busca um prêmio por ID
func (r *RewardRepository) GetByID(id uuid.UUID) (*models.Reward, error) {
	query := `
		SELECT id, name, description, image, draw_date, completed, created_at, updated_at
		FROM rewards WHERE id = $1
	`

	reward := &models.Reward{}
	err := r.db.QueryRow(query, id).Scan(
		&reward.ID, &reward.Name, &reward.Description, &reward.Image,
		&reward.DrawDate, &reward.Completed, &reward.CreatedAt, &reward.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("prêmio não encontrado")
		}
		return nil, err
	}

	return reward, nil
}

// GetDetailsByID busca os detalhes completos de um prêmio
func (r *RewardRepository) GetDetailsByID(id uuid.UUID) (*models.RewardDetails, error) {
	// Buscar o prêmio básico
	reward, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Buscar imagens adicionais
	var images []string
	imagesQuery := `SELECT image_url FROM reward_images WHERE reward_id = $1 ORDER BY created_at`
	rows, err := r.db.Query(imagesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var image string
		if err := rows.Scan(&image); err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	// Buscar compradores com números
	var buyers []models.BuyerWithNumber
	buyersQuery := `
		SELECT u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at, rb.number
		FROM users u
		INNER JOIN reward_buyers rb ON u.id = rb.user_id
		WHERE rb.reward_id = $1
		ORDER BY rb.number
	`
	buyerRows, err := r.db.Query(buyersQuery, id)
	if err != nil {
		return nil, err
	}
	defer buyerRows.Close()

	for buyerRows.Next() {
		var buyer models.BuyerWithNumber
		var user models.User
		err := buyerRows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Role,
			&user.Active, &user.CreatedAt, &user.UpdatedAt, &buyer.Number)
		if err != nil {
			return nil, err
		}
		buyer.User = models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		buyers = append(buyers, buyer)
	}

	// Buscar preço e quota mínima
	var price float64
	var minQuota int
	detailsQuery := `SELECT price, min_quota FROM reward_details WHERE reward_id = $1`
	err = r.db.QueryRow(detailsQuery, id).Scan(&price, &minQuota)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	rewardDetails := &models.RewardDetails{
		Reward:   *reward,
		Images:   images,
		Price:    price,
		MinQuota: minQuota,
		Buyers:   buyers,
	}

	return rewardDetails, nil
}

// List busca todos os prêmios com paginação
func (r *RewardRepository) List(page, limit int, search string) ([]models.Reward, int, error) {
	offset := (page - 1) * limit

	// Query base
	baseQuery := `FROM rewards WHERE 1=1`
	args := []interface{}{}
	argCount := 1

	// Adicionar filtro de busca se fornecido
	if search != "" {
		baseQuery += fmt.Sprintf(" AND (name ILIKE $%d OR description ILIKE $%d)", argCount, argCount)
		args = append(args, "%"+search+"%")
		argCount++
	}

	// Query para contar total
	countQuery := `SELECT COUNT(*) ` + baseQuery
	var total int
	err := r.db.QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query para buscar dados
	selectQuery := `
		SELECT id, name, description, image, draw_date, completed, created_at, updated_at
		` + baseQuery + ` ORDER BY created_at DESC LIMIT $` + fmt.Sprintf("%d", argCount) + ` OFFSET $` + fmt.Sprintf("%d", argCount+1)

	args = append(args, limit, offset)

	rows, err := r.db.Query(selectQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rewards []models.Reward
	for rows.Next() {
		var reward models.Reward
		err := rows.Scan(
			&reward.ID, &reward.Name, &reward.Description, &reward.Image,
			&reward.DrawDate, &reward.Completed, &reward.CreatedAt, &reward.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		rewards = append(rewards, reward)
	}

	return rewards, total, nil
}

// Update atualiza um prêmio
func (r *RewardRepository) Update(id uuid.UUID, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	// Construir query dinamicamente
	setParts := []string{}
	args := []interface{}{}
	argCount := 1

	for field, value := range updates {
		if value != nil {
			setParts = append(setParts, fmt.Sprintf("%s = $%d", field, argCount))
			args = append(args, value)
			argCount++
		}
	}

	// Adicionar updated_at
	setParts = append(setParts, fmt.Sprintf("updated_at = $%d", argCount))
	args = append(args, time.Now())
	argCount++

	// Adicionar ID para WHERE
	args = append(args, id)

	query := fmt.Sprintf("UPDATE rewards SET %s WHERE id = $%d", strings.Join(setParts, ", "), argCount)

	_, err := r.db.Exec(query, args...)
	return err
}

// Delete remove um prêmio
func (r *RewardRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM rewards WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// AddBuyer adiciona um comprador a um prêmio
func (r *RewardRepository) AddBuyer(rewardID, userID uuid.UUID, number int) error {
	query := `INSERT INTO reward_buyers (reward_id, user_id, number, created_at) VALUES ($1, $2, $3, NOW())`
	_, err := r.db.Exec(query, rewardID, userID, number)
	return err
}

// RemoveBuyer remove um comprador de um prêmio
func (r *RewardRepository) RemoveBuyer(rewardID, userID uuid.UUID) error {
	query := `DELETE FROM reward_buyers WHERE reward_id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, rewardID, userID)
	return err
}

// GetBuyers busca todos os compradores de um prêmio com números
func (r *RewardRepository) GetBuyers(rewardID uuid.UUID) ([]models.BuyerWithNumber, error) {
	query := `
		SELECT u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at, rb.number
		FROM users u
		INNER JOIN reward_buyers rb ON u.id = rb.user_id
		WHERE rb.reward_id = $1
		ORDER BY rb.number
	`

	rows, err := r.db.Query(query, rewardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buyers []models.BuyerWithNumber
	for rows.Next() {
		var buyer models.BuyerWithNumber
		var user models.User
		err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Role,
			&user.Active, &user.CreatedAt, &user.UpdatedAt, &buyer.Number)
		if err != nil {
			return nil, err
		}
		buyer.User = models.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			Active:    user.Active,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		buyers = append(buyers, buyer)
	}

	return buyers, nil
}

// GetAvailableNumbers busca os números disponíveis para um prêmio
func (r *RewardRepository) GetAvailableNumbers(rewardID uuid.UUID) ([]int, error) {
	// Buscar números já comprados
	query := `SELECT number FROM reward_buyers WHERE reward_id = $1 ORDER BY number`
	rows, err := r.db.Query(query, rewardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var takenNumbers []int
	for rows.Next() {
		var number int
		if err := rows.Scan(&number); err != nil {
			return nil, err
		}
		takenNumbers = append(takenNumbers, number)
	}

	// Buscar detalhes do prêmio para saber a quota mínima
	var minQuota int
	detailsQuery := `SELECT min_quota FROM reward_details WHERE reward_id = $1`
	err = r.db.QueryRow(detailsQuery, rewardID).Scan(&minQuota)
	if err != nil {
		return nil, err
	}

	// Gerar números disponíveis (de 1 até min_quota, excluindo os já comprados)
	availableNumbers := make([]int, 0)
	takenMap := make(map[int]bool)
	for _, num := range takenNumbers {
		takenMap[num] = true
	}

	for i := 1; i <= minQuota; i++ {
		if !takenMap[i] {
			availableNumbers = append(availableNumbers, i)
		}
	}

	return availableNumbers, nil
}

// BuyNumbers compra uma quantidade específica de números para um usuário
func (r *RewardRepository) BuyNumbers(rewardID, userID uuid.UUID, quantity int) ([]int, error) {
	// Iniciar transação
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Buscar números disponíveis
	availableNumbers, err := r.GetAvailableNumbers(rewardID)
	if err != nil {
		return nil, err
	}

	if len(availableNumbers) < quantity {
		return nil, errors.New("números insuficientes disponíveis")
	}

	// Pegar os primeiros números disponíveis
	numbersToBuy := availableNumbers[:quantity]

	// Inserir cada número comprado
	insertQuery := `INSERT INTO reward_buyers (reward_id, user_id, number, created_at) VALUES ($1, $2, $3, NOW())`

	for _, number := range numbersToBuy {
		_, err = tx.Exec(insertQuery, rewardID, userID, number)
		if err != nil {
			return nil, err
		}
	}

	// Commit da transação
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return numbersToBuy, nil
}
