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
		INSERT INTO rewards (id, owner_id, name, description, image, draw_date, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err = tx.Exec(rewardQuery,
		reward.ID,
		reward.OwnerID,
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
		SELECT id, owner_id, name, description, image, draw_date, completed, created_at, updated_at
		FROM rewards WHERE id = $1
	`

	reward := &models.Reward{}
	err := r.db.QueryRow(query, id).Scan(
		&reward.ID, &reward.OwnerID, &reward.Name, &reward.Description, &reward.Image,
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

	// Buscar compradores com quantidade de números
	var buyers []models.BuyerWithNumber
	buyersQuery := `
		SELECT u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at, count(rb.number) as total_numbers
		FROM users u
		INNER JOIN reward_buyers rb ON u.id = rb.user_id
		WHERE rb.reward_id = $1
		GROUP BY u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at
		ORDER BY total_numbers DESC
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
			&user.Active, &user.CreatedAt, &user.UpdatedAt, &buyer.TotalNumbers)
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
		SELECT id, owner_id, name, description, image, draw_date, completed, created_at, updated_at
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
			&reward.ID, &reward.OwnerID, &reward.Name, &reward.Description, &reward.Image,
			&reward.DrawDate, &reward.Completed, &reward.CreatedAt, &reward.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		rewards = append(rewards, reward)
	}

	return rewards, total, nil
}

// ListByOwner busca todos os prêmios de um dono específico com paginação
func (r *RewardRepository) ListByOwner(ownerID uuid.UUID, page, limit int) ([]models.Reward, int, error) {
	offset := (page - 1) * limit

	countQuery := `SELECT COUNT(*) FROM rewards WHERE owner_id = $1`
	var total int
	err := r.db.QueryRow(countQuery, ownerID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	selectQuery := `SELECT id, owner_id, name, description, image, draw_date, completed, created_at, updated_at FROM rewards WHERE owner_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(selectQuery, ownerID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rewards []models.Reward
	for rows.Next() {
		var reward models.Reward
		err := rows.Scan(
			&reward.ID, &reward.OwnerID, &reward.Name, &reward.Description, &reward.Image,
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

// UpdateDetails atualiza os detalhes de um prêmio (price, min_quota, images)
func (r *RewardRepository) UpdateDetails(rewardID uuid.UUID, price *float64, minQuota *int, images []string) error {
	// Iniciar transação
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Atualizar price e min_quota se fornecidos
	if price != nil || minQuota != nil {
		updateFields := []string{}
		args := []interface{}{}
		argCount := 1

		if price != nil {
			updateFields = append(updateFields, fmt.Sprintf("price = $%d", argCount))
			args = append(args, *price)
			argCount++
		}
		if minQuota != nil {
			updateFields = append(updateFields, fmt.Sprintf("min_quota = $%d", argCount))
			args = append(args, *minQuota)
			argCount++
		}

		updateFields = append(updateFields, fmt.Sprintf("updated_at = $%d", argCount))
		args = append(args, time.Now())
		argCount++

		args = append(args, rewardID)

		detailsQuery := fmt.Sprintf(`
			UPDATE reward_details 
			SET %s 
			WHERE reward_id = $%d
		`, strings.Join(updateFields, ", "), argCount)

		_, err = tx.Exec(detailsQuery, args...)
		if err != nil {
			return err
		}
	}

	// Atualizar imagens se fornecidas
	if len(images) > 0 {
		// Remover imagens existentes
		deleteQuery := `DELETE FROM reward_images WHERE reward_id = $1`
		_, err = tx.Exec(deleteQuery, rewardID)
		if err != nil {
			return err
		}

		// Inserir novas imagens
		imagesQuery := `
			INSERT INTO reward_images (reward_id, image_url, created_at)
			VALUES ($1, $2, $3)
		`

		for _, image := range images {
			if image != "" { // Só inserir se não for vazio
				_, err = tx.Exec(imagesQuery, rewardID, image, time.Now())
				if err != nil {
					return err
				}
			}
		}
	}

	// Commit da transação
	return tx.Commit()
}

// Delete remove um prêmio
func (r *RewardRepository) Delete(id uuid.UUID) error {
	// Iniciar transação
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Deletar registros filhos na ordem correta (evitar problemas de chave estrangeira)

	// 1. Deletar compradores (reward_buyers)
	deleteBuyersQuery := `DELETE FROM reward_buyers WHERE reward_id = $1`
	_, err = tx.Exec(deleteBuyersQuery, id)
	if err != nil {
		return err
	}

	// 2. Deletar imagens (reward_images)
	deleteImagesQuery := `DELETE FROM reward_images WHERE reward_id = $1`
	_, err = tx.Exec(deleteImagesQuery, id)
	if err != nil {
		return err
	}

	// 3. Deletar detalhes (reward_details)
	deleteDetailsQuery := `DELETE FROM reward_details WHERE reward_id = $1`
	_, err = tx.Exec(deleteDetailsQuery, id)
	if err != nil {
		return err
	}

	// 4. Deletar o prêmio principal (rewards)
	deleteRewardQuery := `DELETE FROM rewards WHERE id = $1`
	_, err = tx.Exec(deleteRewardQuery, id)
	if err != nil {
		return err
	}

	// Commit da transação
	return tx.Commit()
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

// GetBuyers busca todos os compradores de um prêmio com quantidade de números
func (r *RewardRepository) GetBuyers(rewardID uuid.UUID) ([]models.BuyerWithNumber, error) {
	query := `
		SELECT u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at, count(rb.number) as total_numbers
		FROM users u
		INNER JOIN reward_buyers rb ON u.id = rb.user_id
		WHERE rb.reward_id = $1
		GROUP BY u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at
		ORDER BY total_numbers DESC
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
			&user.Active, &user.CreatedAt, &user.UpdatedAt, &buyer.TotalNumbers)
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

// GetUserNumbers busca os números específicos de um usuário em um prêmio
func (r *RewardRepository) GetUserNumbers(rewardID, userID uuid.UUID) ([]int, error) {
	query := `
		SELECT number
		FROM reward_buyers
		WHERE reward_id = $1 AND user_id = $2
		ORDER BY number
	`

	rows, err := r.db.Query(query, rewardID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []int
	for rows.Next() {
		var number int
		if err := rows.Scan(&number); err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

// GetAvailableNumbers retorna apenas o próximo número disponível para compra
func (r *RewardRepository) GetMinNumber(rewardID uuid.UUID) (int, error) {
	// Buscar o maior número já comprado
	query := `SELECT COALESCE(MAX(number), 0) FROM reward_buyers WHERE reward_id = $1`
	var maxNumber int
	err := r.db.QueryRow(query, rewardID).Scan(&maxNumber)
	if err != nil {
		return 1, err
	}

	// O próximo número disponível é o maior + 1
	nextNumber := maxNumber + 1
	return nextNumber, nil
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
	minNumber, err := r.GetMinNumber(rewardID)
	if err != nil {
		return nil, err
	}

	// quero os numeros de minNumber até minNumber + quantity
	numbersToBuy := make([]int, quantity)
	for i := 0; i < quantity; i++ {
		numbersToBuy[i] = minNumber + i
	}

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

// GetUserPurchases busca todas as compras de um usuário
func (r *RewardRepository) GetUserPurchases(userID uuid.UUID, page, limit int) ([]models.Purchase, int, error) {
	offset := (page - 1) * limit

	// Query para contar total
	countQuery := `
		SELECT COUNT(DISTINCT rb.reward_id)
		FROM reward_buyers rb
		WHERE rb.user_id = $1
	`
	var total int
	err := r.db.QueryRow(countQuery, userID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query para buscar compras com paginação
	query := `
		SELECT 
			r.id as reward_id,
			r.name as reward_name,
			r.image as reward_image,
			MIN(rb.created_at) as purchase_date,
			COUNT(rb.number) as total_numbers,
			rd.price as price_per_number,
			r.completed as reward_completed
		FROM reward_buyers rb
		INNER JOIN rewards r ON rb.reward_id = r.id
		LEFT JOIN reward_details rd ON r.id = rd.reward_id
		WHERE rb.user_id = $1
		GROUP BY r.id, r.name, r.image, rd.price, r.completed
		ORDER BY purchase_date DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var purchases []models.Purchase
	purchaseID := 1

	for rows.Next() {
		var purchase models.Purchase
		var rewardID uuid.UUID
		var rewardName, rewardImage string
		var purchaseDate time.Time
		var totalNumbers int
		var pricePerNumber sql.NullFloat64
		var rewardCompleted bool

		err := rows.Scan(
			&rewardID, &rewardName, &rewardImage, &purchaseDate,
			&totalNumbers, &pricePerNumber, &rewardCompleted)
		if err != nil {
			return nil, 0, err
		}

		// Buscar os números específicos desta compra
		numbers, err := r.GetUserNumbers(rewardID, userID)
		if err != nil {
			return nil, 0, err
		}

		// Calcular valor total
		totalAmount := 0.0
		if pricePerNumber.Valid {
			totalAmount = float64(totalNumbers) * pricePerNumber.Float64
		}

		// Determinar status
		status := "active"
		if rewardCompleted {
			status = "completed"
		}

		purchase = models.Purchase{
			ID:           purchaseID,
			RewardID:     rewardID,
			RewardName:   rewardName,
			RewardImage:  rewardImage,
			Numbers:      numbers,
			PurchaseDate: purchaseDate,
			TotalAmount:  totalAmount,
			Status:       status,
		}

		purchases = append(purchases, purchase)
		purchaseID++
	}

	return purchases, total, nil
}
