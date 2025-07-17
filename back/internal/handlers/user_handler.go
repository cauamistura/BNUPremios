package handlers

import (
	"net/http"

	"github.com/cauamistura/BNUPremios/internal/models"
	"github.com/cauamistura/BNUPremios/internal/services"
	"github.com/gin-gonic/gin"
)

// UserHandler implementa os handlers HTTP para usuários
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler cria uma nova instância do handler de usuários
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetByID godoc
// @Summary Buscar usuário por ID
// @Description Busca um usuário específico pelo ID (requer autenticação)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do usuário"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.userService.GetByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "ID inválido" {
			status = http.StatusBadRequest
		} else if err.Error() == "usuário não encontrado" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// List godoc
// @Summary Listar usuários
// @Description Lista todos os usuários com paginação (requer autenticação)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Número da página (padrão: 1)"
// @Param limit query int false "Limite por página (padrão: 10, máximo: 100)"
// @Success 200 {object} models.UserListResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func (h *UserHandler) List(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	users, err := h.userService.List(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// Update godoc
// @Summary Atualizar usuário
// @Description Atualiza os dados de um usuário específico (requer autenticação)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do usuário"
// @Param user body models.UpdateUserRequest true "Dados para atualização"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var updateReq models.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	user, err := h.userService.Update(id, &updateReq)
	if err != nil {
		status := http.StatusInternalServerError
		switch err.Error() {
		case "ID inválido":
			status = http.StatusBadRequest
		case "usuário não encontrado":
			status = http.StatusNotFound
		case "email já está em uso":
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Deletar usuário
// @Description Remove um usuário do sistema (requer autenticação)
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID do usuário"
// @Success 204 "Usuário deletado com sucesso"
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.userService.Delete(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "ID inválido" {
			status = http.StatusBadRequest
		} else if err.Error() == "usuário não encontrado" {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// Login godoc
// @Summary Login de usuário
// @Description Autentica um usuário e retorna um token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Credenciais de login"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var loginReq models.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	loginResponse, err := h.userService.Login(&loginReq)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "credenciais inválidas" || err.Error() == "usuário inativo" {
			status = http.StatusUnauthorized
		}
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

// Register godoc
// @Summary Registrar novo usuário
// @Description Registra um novo usuário no sistema
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "Dados do usuário"
// @Success 201 {object} models.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var registerReq models.RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"details": err.Error(),
		})
		return
	}

	userResponse, err := h.userService.Register(&registerReq)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "email já está em uso" {
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, userResponse)
}
