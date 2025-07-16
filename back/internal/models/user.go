package models

import (
	"time"

	"github.com/google/uuid"
)

// User representa um usuário no sistema
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required,email"`
	Password  string    `json:"password,omitempty" db:"password" binding:"required,min=6"`
	Role      string    `json:"role" db:"role"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// UserResponse representa a resposta de um usuário (sem senha)
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// LoginRequest representa a requisição de login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest representa a requisição de registro
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UpdateUserRequest representa a requisição de atualização de usuário
type UpdateUserRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email" binding:"omitempty,email"`
	Role   string `json:"role"`
	Active *bool  `json:"active"`
}

// LoginResponse representa a resposta de login
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// Pagination representa a paginação
type Pagination struct {
	Page    int  `json:"page"`
	Limit   int  `json:"limit"`
	Total   int  `json:"total"`
	Pages   int  `json:"pages"`
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
}

// UserListResponse representa a resposta da listagem de usuários
type UserListResponse struct {
	Users      []UserResponse `json:"users"`
	Pagination Pagination     `json:"pagination"`
}

// BuyerWithNumber representa um comprador com o número comprado
type BuyerWithNumber struct {
	User   UserResponse `json:"user"`
	Number int          `json:"number"`
}

// BuyerListResponse representa a resposta da listagem de compradores
type BuyerListResponse struct {
	Buyers     []BuyerWithNumber `json:"buyers"`
	Pagination Pagination        `json:"pagination"`
}
