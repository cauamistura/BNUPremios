package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/cauamistura/BNUPremios/internal/models"
	"github.com/cauamistura/BNUPremios/internal/repository"
	"github.com/google/uuid"
)

type RewardService struct {
	rewardRepo *repository.RewardRepository
}

func NewRewardService(rewardRepo *repository.RewardRepository) *RewardService {
	return &RewardService{rewardRepo: rewardRepo}
}

// Create cria um novo prêmio
func (s *RewardService) Create(req *models.CreateRewardRequest, ownerID uuid.UUID) (*models.RewardResponse, error) {
	reward := &models.Reward{
		ID:          uuid.New(),
		OwnerID:     ownerID,
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		DrawDate:    req.DrawDate,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.rewardRepo.Create(reward, req.Price, req.MinQuota, req.Images); err != nil {
		return nil, fmt.Errorf("erro ao criar prêmio: %w", err)
	}

	return s.toRewardResponse(reward), nil
}

// GetByID busca um prêmio por ID
func (s *RewardService) GetByID(id uuid.UUID) (*models.RewardResponse, error) {
	reward, err := s.rewardRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.toRewardResponse(reward), nil
}

// GetDetailsByID busca os detalhes completos de um prêmio
func (s *RewardService) GetDetailsByID(id uuid.UUID) (*models.RewardDetailsResponse, error) {
	rewardDetails, err := s.rewardRepo.GetDetailsByID(id)
	if err != nil {
		return nil, err
	}

	return s.toRewardDetailsResponse(rewardDetails), nil
}

// GetDetailsByIDWithoutBuyers busca os detalhes completos de um prêmio sem compradores
func (s *RewardService) GetDetailsByIDWithoutBuyers(id uuid.UUID) (*models.RewardDetailsWithoutBuyersResponse, error) {
	rewardDetails, err := s.rewardRepo.GetDetailsByID(id)
	if err != nil {
		return nil, err
	}

	return s.ToRewardDetailsWithoutBuyersResponse(rewardDetails), nil
}

// List busca todos os prêmios com paginação
func (s *RewardService) List(page, limit int, search string) (*models.RewardListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	rewards, total, err := s.rewardRepo.List(page, limit, search)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar prêmios: %w", err)
	}

	// Converter para response
	var rewardResponses []models.RewardResponse
	for _, reward := range rewards {
		rewardResponses = append(rewardResponses, *s.toRewardResponse(&reward))
	}

	// Calcular paginação
	pages := (total + limit - 1) / limit
	hasNext := page < pages
	hasPrev := page > 1

	pagination := models.Pagination{
		Page:    page,
		Limit:   limit,
		Total:   total,
		Pages:   pages,
		HasNext: hasNext,
		HasPrev: hasPrev,
	}

	return &models.RewardListResponse{
		Rewards:    rewardResponses,
		Pagination: pagination,
	}, nil
}

// Update atualiza um prêmio
func (s *RewardService) Update(id uuid.UUID, req *models.UpdateRewardRequest) (*models.RewardResponse, error) {
	// Verificar se o prêmio existe
	_, err := s.rewardRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Construir map de atualizações
	updates := make(map[string]interface{})
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = *req.Description
	}
	if req.Image != nil {
		updates["image"] = *req.Image
	}
	if req.DrawDate != nil {
		updates["draw_date"] = *req.DrawDate
	}
	if req.Completed != nil {
		updates["completed"] = *req.Completed
	}

	// Atualizar no banco se houver mudanças
	if len(updates) > 0 {
		if err := s.rewardRepo.Update(id, updates); err != nil {
			return nil, fmt.Errorf("erro ao atualizar prêmio: %w", err)
		}
	}

	// Atualizar detalhes do prêmio (price, min_quota, images) se fornecidos
	if req.Price != nil || req.MinQuota != nil || len(req.Images) > 0 {
		if err := s.rewardRepo.UpdateDetails(id, req.Price, req.MinQuota, req.Images); err != nil {
			return nil, fmt.Errorf("erro ao atualizar detalhes do prêmio: %w", err)
		}
	}

	// Buscar prêmio atualizado
	updatedReward, err := s.rewardRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.toRewardResponse(updatedReward), nil
}

// Delete remove um prêmio
func (s *RewardService) Delete(id uuid.UUID) error {
	// Verificar se o prêmio existe
	_, err := s.rewardRepo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.rewardRepo.Delete(id); err != nil {
		return fmt.Errorf("erro ao deletar prêmio: %w", err)
	}

	return nil
}

// AddBuyer adiciona um comprador a um prêmio
func (s *RewardService) AddBuyer(rewardID, userID uuid.UUID, number int) error {
	// Verificar se o prêmio existe
	_, err := s.rewardRepo.GetByID(rewardID)
	if err != nil {
		return errors.New("prêmio não encontrado")
	}

	if err := s.rewardRepo.AddBuyer(rewardID, userID, number); err != nil {
		return fmt.Errorf("erro ao adicionar comprador: %w", err)
	}

	return nil
}

// RemoveBuyer remove um comprador de um prêmio
func (s *RewardService) RemoveBuyer(rewardID, userID uuid.UUID) error {
	if err := s.rewardRepo.RemoveBuyer(rewardID, userID); err != nil {
		return fmt.Errorf("erro ao remover comprador: %w", err)
	}

	return nil
}

// GetBuyers busca todos os compradores de um prêmio
func (s *RewardService) GetBuyers(rewardID uuid.UUID) ([]models.BuyerWithNumber, error) {
	buyers, err := s.rewardRepo.GetBuyers(rewardID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar compradores: %w", err)
	}

	return buyers, nil
}

// BuyNumbers compra uma quantidade específica de números para um usuário
func (s *RewardService) BuyNumbers(rewardID, userID uuid.UUID, quantity int) ([]int, error) {
	// Verificar se o prêmio existe
	_, err := s.rewardRepo.GetByID(rewardID)
	if err != nil {
		return nil, errors.New("prêmio não encontrado")
	}

	if quantity <= 0 {
		return nil, errors.New("quantidade deve ser maior que zero")
	}

	numbers, err := s.rewardRepo.BuyNumbers(rewardID, userID, quantity)
	if err != nil {
		return nil, fmt.Errorf("erro ao comprar números: %w", err)
	}

	return numbers, nil
}

// GetUserNumbers busca os números específicos de um usuário em um prêmio
func (s *RewardService) GetUserNumbers(rewardID, userID uuid.UUID) ([]int, error) {
	// Verificar se o prêmio existe
	_, err := s.rewardRepo.GetByID(rewardID)
	if err != nil {
		return nil, errors.New("prêmio não encontrado")
	}

	numbers, err := s.rewardRepo.GetUserNumbers(rewardID, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar números do usuário: %w", err)
	}

	return numbers, nil
}

// GetUserPurchases busca todas as compras de um usuário
func (s *RewardService) GetUserPurchases(userID uuid.UUID, page, limit int) (*models.PurchaseListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	purchases, total, err := s.rewardRepo.GetUserPurchases(userID, page, limit)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar compras do usuário: %w", err)
	}

	// Calcular paginação
	pages := (total + limit - 1) / limit
	hasNext := page < pages
	hasPrev := page > 1

	pagination := models.Pagination{
		Page:    page,
		Limit:   limit,
		Total:   total,
		Pages:   pages,
		HasNext: hasNext,
		HasPrev: hasPrev,
	}

	return &models.PurchaseListResponse{
		Purchases:  purchases,
		Pagination: pagination,
	}, nil
}

// ListByOwner busca todos os prêmios de um dono específico com paginação
func (s *RewardService) ListByOwner(ownerID uuid.UUID, page, limit int) (*models.RewardListResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	rewards, total, err := s.rewardRepo.ListByOwner(ownerID, page, limit)
	if err != nil {
		return nil, err
	}

	var rewardResponses []models.RewardResponse
	for _, reward := range rewards {
		rewardResponses = append(rewardResponses, *s.toRewardResponse(&reward))
	}

	pages := (total + limit - 1) / limit
	hasNext := page < pages
	hasPrev := page > 1

	pagination := models.Pagination{
		Page:    page,
		Limit:   limit,
		Total:   total,
		Pages:   pages,
		HasNext: hasNext,
		HasPrev: hasPrev,
	}

	return &models.RewardListResponse{
		Rewards:    rewardResponses,
		Pagination: pagination,
	}, nil
}

// toRewardResponse converte Reward para RewardResponse
func (s *RewardService) toRewardResponse(reward *models.Reward) *models.RewardResponse {
	return &models.RewardResponse{
		ID:          reward.ID,
		OwnerID:     reward.OwnerID,
		Name:        reward.Name,
		Description: reward.Description,
		Image:       reward.Image,
		DrawDate:    reward.DrawDate,
		Completed:   reward.Completed,
		CreatedAt:   reward.CreatedAt,
		UpdatedAt:   reward.UpdatedAt,
	}
}

// toRewardDetailsResponse converte RewardDetails para RewardDetailsResponse
func (s *RewardService) toRewardDetailsResponse(rewardDetails *models.RewardDetails) *models.RewardDetailsResponse {
	rewardResponse := s.toRewardResponse(&rewardDetails.Reward)

	return &models.RewardDetailsResponse{
		RewardResponse: *rewardResponse,
		Images:         rewardDetails.Images,
		Price:          rewardDetails.Price,
		MinQuota:       rewardDetails.MinQuota,
		Buyers:         rewardDetails.Buyers,
	}
}

// ToRewardDetailsWithoutBuyersResponse converte RewardDetails para RewardDetailsWithoutBuyersResponse
func (s *RewardService) ToRewardDetailsWithoutBuyersResponse(rewardDetails *models.RewardDetails) *models.RewardDetailsWithoutBuyersResponse {
	rewardResponse := s.toRewardResponse(&rewardDetails.Reward)

	return &models.RewardDetailsWithoutBuyersResponse{
		RewardResponse: *rewardResponse,
		Images:         rewardDetails.Images,
		Price:          rewardDetails.Price,
		MinQuota:       rewardDetails.MinQuota,
	}
}

// toUserResponse converte User para UserResponse
func (s *RewardService) toUserResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
