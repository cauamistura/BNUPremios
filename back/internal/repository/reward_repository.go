package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
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
		SELECT id, owner_id, name, description, image, draw_date, completed, winner_number, drawn_at, created_at, updated_at
		FROM rewards
		WHERE id = $1
	`

	var reward models.Reward
	err := r.db.QueryRow(query, id).Scan(
		&reward.ID, &reward.OwnerID, &reward.Name, &reward.Description,
		&reward.Image, &reward.DrawDate, &reward.Completed, &reward.WinnerNumber, &reward.DrawnAt,
		&reward.CreatedAt, &reward.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &reward, nil
}

// GetDetailsByID busca os detalhes completos de um prêmio
func (r *RewardRepository) GetDetailsByID(id uuid.UUID) (*models.RewardDetails, error) {
	// Buscar dados básicos do prêmio
	reward, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Buscar detalhes (price e min_quota)
	var price float64
	var minQuota int
	detailsQuery := `SELECT price, min_quota FROM reward_details WHERE reward_id = $1`
	err = r.db.QueryRow(detailsQuery, id).Scan(&price, &minQuota)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Buscar imagens
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

	// Buscar compradores
	buyers, err := r.GetBuyers(id)
	if err != nil {
		return nil, err
	}

	// Buscar ganhador se o prêmio foi sorteado
	var winnerUser *models.UserResponse
	if reward.WinnerNumber != nil {
		winner, err := r.GetWinnerByNumber(id, *reward.WinnerNumber)
		if err == nil {
			winnerUser = &models.UserResponse{
				ID:        winner.ID,
				Name:      winner.Name,
				Email:     winner.Email,
				Role:      winner.Role,
				Active:    winner.Active,
				CreatedAt: winner.CreatedAt,
				UpdatedAt: winner.UpdatedAt,
			}
		}
	}

	return &models.RewardDetails{
		Reward:     *reward,
		Images:     images,
		Price:      price,
		MinQuota:   minQuota,
		Buyers:     buyers,
		WinnerUser: winnerUser,
	}, nil
}

// List busca todos os prêmios com paginação
func (r *RewardRepository) List(page, limit int, search string) ([]models.Reward, int, error) {
	offset := (page - 1) * limit

	// Query para contar total
	countQuery := `SELECT COUNT(*) FROM rewards`
	var total int
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query para buscar prêmios
	query := `
		SELECT id, owner_id, name, description, image, draw_date, completed, winner_number, drawn_at, created_at, updated_at
		FROM rewards
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rewards []models.Reward
	for rows.Next() {
		var reward models.Reward
		err := rows.Scan(
			&reward.ID, &reward.OwnerID, &reward.Name, &reward.Description,
			&reward.Image, &reward.DrawDate, &reward.Completed, &reward.WinnerNumber, &reward.DrawnAt,
			&reward.CreatedAt, &reward.UpdatedAt,
		)
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

	// Query para contar total
	countQuery := `SELECT COUNT(*) FROM rewards WHERE owner_id = $1`
	var total int
	err := r.db.QueryRow(countQuery, ownerID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	// Query para buscar prêmios
	query := `
		SELECT id, owner_id, name, description, image, draw_date, completed, winner_number, drawn_at, created_at, updated_at
		FROM rewards
		WHERE owner_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.Query(query, ownerID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var rewards []models.Reward
	for rows.Next() {
		var reward models.Reward
		err := rows.Scan(
			&reward.ID, &reward.OwnerID, &reward.Name, &reward.Description,
			&reward.Image, &reward.DrawDate, &reward.Completed, &reward.WinnerNumber, &reward.DrawnAt,
			&reward.CreatedAt, &reward.UpdatedAt,
		)
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

	// Verificar se o prêmio está completado
	var completed bool
	checkQuery := `SELECT completed FROM rewards WHERE id = $1`
	err = tx.QueryRow(checkQuery, rewardID).Scan(&completed)
	if err != nil {
		return nil, err
	}

	if completed {
		return nil, errors.New("não é possível comprar números de um prêmio já completado")
	}

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

// DrawReward realiza o sorteio de um prêmio
func (r *RewardRepository) DrawReward(rewardID uuid.UUID) (*models.DrawRewardResponse, error) {
	// Iniciar transação
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Verificar se o prêmio já foi sorteado
	var winnerNumber *int
	var drawnAt *time.Time
	checkQuery := `SELECT winner_number, drawn_at FROM rewards WHERE id = $1`
	err = tx.QueryRow(checkQuery, rewardID).Scan(&winnerNumber, &drawnAt)
	if err != nil {
		return nil, err
	}

	if winnerNumber != nil {
		return nil, errors.New("prêmio já foi sorteado")
	}

	// Buscar todos os números comprados para este prêmio
	numbersQuery := `
		SELECT number, user_id 
		FROM reward_buyers 
		WHERE reward_id = $1 
		ORDER BY number
	`
	rows, err := tx.Query(numbersQuery, rewardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var numbers []int
	var userIDs []uuid.UUID
	numberToUser := make(map[int]uuid.UUID)

	for rows.Next() {
		var number int
		var userID uuid.UUID
		if err := rows.Scan(&number, &userID); err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
		userIDs = append(userIDs, userID)
		numberToUser[number] = userID
	}

	if len(numbers) == 0 {
		return nil, errors.New("nenhum número foi comprado para este prêmio")
	}

	// Realizar sorteio aleatório
	rand.Seed(time.Now().UnixNano())
	winnerIndex := rand.Intn(len(numbers))
	winnerNumber = &numbers[winnerIndex]
	winnerUserID := numberToUser[*winnerNumber]

	// Atualizar o prêmio com o número vencedor e marcar como completado
	now := time.Now()
	updateQuery := `
		UPDATE rewards 
		SET winner_number = $1, drawn_at = $2, completed = true, updated_at = $3 
		WHERE id = $4
	`
	_, err = tx.Exec(updateQuery, *winnerNumber, now, now, rewardID)
	if err != nil {
		return nil, err
	}

	// Buscar informações do usuário vencedor
	var winnerUser models.User
	userQuery := `
		SELECT id, name, email, role, active, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`
	err = tx.QueryRow(userQuery, winnerUserID).Scan(
		&winnerUser.ID, &winnerUser.Name, &winnerUser.Email,
		&winnerUser.Role, &winnerUser.Active, &winnerUser.CreatedAt, &winnerUser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// Commit da transação
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	// Converter para response
	winnerUserResponse := models.UserResponse{
		ID:        winnerUser.ID,
		Name:      winnerUser.Name,
		Email:     winnerUser.Email,
		Role:      winnerUser.Role,
		Active:    winnerUser.Active,
		CreatedAt: winnerUser.CreatedAt,
		UpdatedAt: winnerUser.UpdatedAt,
	}

	return &models.DrawRewardResponse{
		RewardID:     rewardID,
		WinnerNumber: *winnerNumber,
		WinnerUser:   &winnerUserResponse,
		DrawnAt:      now,
		Message:      fmt.Sprintf("Sorteio realizado! Número vencedor: %d. Prêmio marcado como completado.", *winnerNumber),
	}, nil
}

// GetWinnerByNumber busca o usuário que comprou um número específico
func (r *RewardRepository) GetWinnerByNumber(rewardID uuid.UUID, number int) (*models.User, error) {
	query := `
		SELECT u.id, u.name, u.email, u.role, u.active, u.created_at, u.updated_at
		FROM users u
		INNER JOIN reward_buyers rb ON u.id = rb.user_id
		WHERE rb.reward_id = $1 AND rb.number = $2
	`

	var user models.User
	err := r.db.QueryRow(query, rewardID, number).Scan(
		&user.ID, &user.Name, &user.Email,
		&user.Role, &user.Active, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// IsRewardDrawn verifica se um prêmio já foi sorteado
func (r *RewardRepository) IsRewardDrawn(rewardID uuid.UUID) (bool, error) {
	var winnerNumber *int
	query := `SELECT winner_number FROM rewards WHERE id = $1`
	err := r.db.QueryRow(query, rewardID).Scan(&winnerNumber)
	if err != nil {
		return false, err
	}
	return winnerNumber != nil, nil
}
