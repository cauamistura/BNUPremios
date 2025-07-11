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
func (s *RewardService) Create(req *models.CreateRewardRequest) (*models.RewardResponse, error) {
	reward := &models.Reward{
		ID:          uuid.New(),
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
	existingReward, err := s.rewardRepo.GetByID(id)
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

	if len(updates) == 0 {
		return s.toRewardResponse(existingReward), nil
	}

	// Atualizar no banco
	if err := s.rewardRepo.Update(id, updates); err != nil {
		return nil, fmt.Errorf("erro ao atualizar prêmio: %w", err)
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
func (s *RewardService) AddBuyer(rewardID, userID uuid.UUID) error {
	// Verificar se o prêmio existe
	_, err := s.rewardRepo.GetByID(rewardID)
	if err != nil {
		return errors.New("prêmio não encontrado")
	}

	if err := s.rewardRepo.AddBuyer(rewardID, userID); err != nil {
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
func (s *RewardService) GetBuyers(rewardID uuid.UUID) ([]models.UserResponse, error) {
	buyers, err := s.rewardRepo.GetBuyers(rewardID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar compradores: %w", err)
	}

	var buyerResponses []models.UserResponse
	for _, buyer := range buyers {
		buyerResponses = append(buyerResponses, *s.toUserResponse(&buyer))
	}

	return buyerResponses, nil
}

// toRewardResponse converte Reward para RewardResponse
func (s *RewardService) toRewardResponse(reward *models.Reward) *models.RewardResponse {
	return &models.RewardResponse{
		ID:          reward.ID,
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
	
	var buyerResponses []models.UserResponse
	for _, buyer := range rewardDetails.Buyers {
		buyerResponses = append(buyerResponses, *s.toUserResponse(&buyer))
	}

	return &models.RewardDetailsResponse{
		RewardResponse: *rewardResponse,
		Images:         rewardDetails.Images,
		Price:          rewardDetails.Price,
		MinQuota:       rewardDetails.MinQuota,
		Buyers:         buyerResponses,
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