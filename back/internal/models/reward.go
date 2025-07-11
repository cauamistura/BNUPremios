package models

import (
	"time"

	"github.com/google/uuid"
)

// Reward representa um prêmio no sistema
type Reward struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description string    `json:"description" db:"description"`
	Image       string    `json:"image" db:"image"`
	DrawDate    time.Time `json:"draw_date" db:"draw_date"`
	Completed   bool      `json:"completed" db:"completed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// RewardDetails representa os detalhes completos de um prêmio
type RewardDetails struct {
	Reward
	Images   []string `json:"images" db:"images"`
	Price    float64  `json:"price" db:"price"`
	MinQuota int      `json:"min_quota" db:"min_quota"`
	Buyers   []User   `json:"buyers" db:"buyers"`
}

// CreateRewardRequest representa a requisição de criação de prêmio
type CreateRewardRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	DrawDate    time.Time `json:"draw_date" binding:"required"`
	Images      []string  `json:"images"`
	Price       float64   `json:"price"`
	MinQuota    int       `json:"min_quota"`
}

// UpdateRewardRequest representa a requisição de atualização de prêmio
type UpdateRewardRequest struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Image       *string    `json:"image"`
	DrawDate    *time.Time `json:"draw_date"`
	Completed   *bool      `json:"completed"`
	Images      []string   `json:"images"`
	Price       *float64   `json:"price"`
	MinQuota    *int       `json:"min_quota"`
}

// RewardResponse representa a resposta de um prêmio
type RewardResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	DrawDate    time.Time `json:"draw_date"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RewardDetailsResponse representa a resposta detalhada de um prêmio
type RewardDetailsResponse struct {
	RewardResponse
	Images   []string       `json:"images"`
	Price    float64        `json:"price"`
	MinQuota int            `json:"min_quota"`
	Buyers   []UserResponse `json:"buyers"`
}

// RewardListResponse representa a resposta da listagem de prêmios
type RewardListResponse struct {
	Rewards    []RewardResponse `json:"rewards"`
	Pagination Pagination       `json:"pagination"`
} 