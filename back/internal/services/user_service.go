package services

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/cauamistura/BNUPremios/internal/models"
	"github.com/cauamistura/BNUPremios/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// UserService implementa a lógica de negócio para usuários
type UserService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

// NewUserService cria uma nova instância do serviço de usuários
func NewUserService(userRepo *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{userRepo: userRepo, jwtSecret: jwtSecret}
}

// Create cria um novo usuário
func (s *UserService) Create(user *models.User) (*models.UserResponse, error) {
	// Verificar se o email já existe
	exists, err := s.userRepo.EmailExists(user.Email)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar email: %w", err)
	}
	if exists {
		return nil, errors.New("email já está em uso")
	}

	// Criptografar senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao criptografar senha: %w", err)
	}
	user.Password = string(hashedPassword)

	// Criar usuário
	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	// Retornar resposta sem senha
	return s.toUserResponse(user), nil
}

// GetByID busca um usuário pelo ID
func (s *UserService) GetByID(id string) (*models.UserResponse, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(user), nil
}

// List busca todos os usuários com paginação
func (s *UserService) List(pageStr, limitStr string) (*models.UserListResponse, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	users, total, err := s.userRepo.List(page, limit)
	if err != nil {
		return nil, err
	}

	// Converter para UserResponse
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, *s.toUserResponse(&user))
	}

	// Calcular paginação
	pages := int(math.Ceil(float64(total) / float64(limit)))
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

	return &models.UserListResponse{
		Users:      userResponses,
		Pagination: pagination,
	}, nil
}

// Update atualiza um usuário
func (s *UserService) Update(id string, updates *models.UpdateUserRequest) (*models.UserResponse, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	// Verificar se o usuário existe
	existingUser, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	// Preparar updates
	updateMap := make(map[string]interface{})
	if updates.Name != "" {
		updateMap["name"] = updates.Name
	}
	if updates.Email != "" {
		// Verificar se o novo email já existe (se for diferente do atual)
		if updates.Email != existingUser.Email {
			exists, err := s.userRepo.EmailExists(updates.Email)
			if err != nil {
				return nil, fmt.Errorf("erro ao verificar email: %w", err)
			}
			if exists {
				return nil, errors.New("email já está em uso")
			}
		}
		updateMap["email"] = updates.Email
	}
	if updates.Role != "" {
		updateMap["role"] = updates.Role
	}
	if updates.Active != nil {
		updateMap["active"] = *updates.Active
	}

	// Atualizar usuário
	if err := s.userRepo.Update(userID, updateMap); err != nil {
		return nil, err
	}

	// Buscar usuário atualizado
	updatedUser, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	return s.toUserResponse(updatedUser), nil
}

// Delete remove um usuário
func (s *UserService) Delete(id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return errors.New("ID inválido")
	}

	return s.userRepo.Delete(userID)
}

// Login autentica um usuário
func (s *UserService) Login(loginReq *models.LoginRequest) (*models.LoginResponse, error) {
	// Buscar usuário por email
	user, err := s.userRepo.GetByEmail(loginReq.Email)
	if err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	// Verificar se o usuário está ativo
	if !user.Active {
		return nil, errors.New("usuário inativo")
	}

	// Verificar senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	// Gerar token JWT real (simples, sem validação de segredo forte por enquanto)
	claims := jwt.MapClaims{
		"sub":     user.ID.String(),
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, errors.New("erro ao gerar token JWT")
	}

	return &models.LoginResponse{
		Token: tokenString,
		User:  *s.toUserResponse(user),
	}, nil
}

// Register registra um novo usuário
func (s *UserService) Register(registerReq *models.RegisterRequest) (*models.UserResponse, error) {
	user := &models.User{
		Name:     registerReq.Name,
		Email:    registerReq.Email,
		Password: registerReq.Password,
		Role:     "user",
		Active:   true,
	}

	return s.Create(user)
}

// toUserResponse converte User para UserResponse
func (s *UserService) toUserResponse(user *models.User) *models.UserResponse {
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
