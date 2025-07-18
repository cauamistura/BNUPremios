package models

import (
	"time"

	"github.com/google/uuid"
)

// Reward representa um prêmio no sistema
type Reward struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	OwnerID      uuid.UUID  `json:"owner_id" db:"owner_id"`
	Name         string     `json:"name" db:"name" binding:"required"`
	Description  string     `json:"description" db:"description"`
	Image        string     `json:"image" db:"image"`
	DrawDate     time.Time  `json:"draw_date" db:"draw_date"`
	Completed    bool       `json:"completed" db:"completed"`
	WinnerNumber *int       `json:"winner_number,omitempty" db:"winner_number"`
	DrawnAt      *time.Time `json:"drawn_at,omitempty" db:"drawn_at"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

// RewardDetails representa os detalhes completos de um prêmio
type RewardDetails struct {
	Reward     Reward            `json:"reward"`
	Images     []string          `json:"images"`
	Price      float64           `json:"price"`
	MinQuota   int               `json:"min_quota"`
	Buyers     []BuyerWithNumber `json:"buyers"`
	WinnerUser *UserResponse     `json:"winner_user,omitempty"`
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
	ID           uuid.UUID  `json:"id"`
	OwnerID      uuid.UUID  `json:"owner_id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Image        string     `json:"image"`
	DrawDate     time.Time  `json:"draw_date"`
	Completed    bool       `json:"completed"`
	WinnerNumber *int       `json:"winner_number,omitempty"`
	DrawnAt      *time.Time `json:"drawn_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// RewardDetailsResponse representa a resposta com detalhes completos de um prêmio
type RewardDetailsResponse struct {
	RewardResponse
	Images     []string          `json:"images"`
	Price      float64           `json:"price"`
	MinQuota   int               `json:"min_quota"`
	Buyers     []BuyerWithNumber `json:"buyers"`
	WinnerUser *UserResponse     `json:"winner_user,omitempty"`
}

// RewardDetailsWithoutBuyersResponse representa a resposta com detalhes de um prêmio sem compradores
type RewardDetailsWithoutBuyersResponse struct {
	RewardResponse
	Images     []string      `json:"images"`
	Price      float64       `json:"price"`
	MinQuota   int           `json:"min_quota"`
	WinnerUser *UserResponse `json:"winner_user,omitempty"`
}

// RewardListResponse representa a resposta da listagem de prêmios
type RewardListResponse struct {
	Rewards    []RewardResponse `json:"rewards"`
	Pagination Pagination       `json:"pagination"`
}

// BuyNumbersRequest representa a requisição para comprar números
type BuyNumbersRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

// DrawRewardRequest representa a requisição para realizar o sorteio
type DrawRewardRequest struct {
	// Pode ser vazio, o sorteio será automático
}

// DrawRewardResponse representa a resposta do sorteio
type DrawRewardResponse struct {
	RewardID     uuid.UUID     `json:"reward_id"`
	WinnerNumber int           `json:"winner_number"`
	WinnerUser   *UserResponse `json:"winner_user,omitempty"`
	DrawnAt      time.Time     `json:"drawn_at"`
	Message      string        `json:"message"`
}
