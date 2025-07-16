package handlers

import (
	"net/http"
	"strconv"

	"github.com/cauamistura/BNUPremios/internal/models"
	"github.com/cauamistura/BNUPremios/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RewardHandler struct {
	rewardService *services.RewardService
}

func NewRewardHandler(rewardService *services.RewardService) *RewardHandler {
	return &RewardHandler{rewardService: rewardService}
}

// Create @Summary Criar prêmio
// @Description Cria um novo prêmio
// @Tags rewards
// @Accept json
// @Produce json
// @Param reward body models.CreateRewardRequest true "Dados do prêmio"
// @Success 201 {object} models.RewardResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards [post]
func (h *RewardHandler) Create(c *gin.Context) {
	var req models.CreateRewardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"message": err.Error(),
		})
		return
	}

	reward, err := h.rewardService.Create(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, reward)
}

// GetByID @Summary Buscar prêmio por ID
// @Description Busca um prêmio específico pelo ID
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Success 200 {object} models.RewardResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id} [get]
func (h *RewardHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	reward, err := h.rewardService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Prêmio não encontrado",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reward)
}

// GetDetailsByID @Summary Buscar detalhes do prêmio por ID
// @Description Busca os detalhes completos de um prêmio específico pelo ID
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Success 200 {object} models.RewardDetailsResponse
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id}/details [get]
func (h *RewardHandler) GetDetailsByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	rewardDetails, err := h.rewardService.GetDetailsByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Prêmio não encontrado",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rewardDetails)
}

// List @Summary Listar prêmios
// @Description Lista todos os prêmios com paginação
// @Tags rewards
// @Accept json
// @Produce json
// @Param page query int false "Página (padrão: 1)"
// @Param limit query int false "Limite por página (padrão: 10, máximo: 100)"
// @Param search query string false "Termo de busca"
// @Success 200 {object} models.RewardListResponse
// @Failure 500 {object} map[string]interface{}
// @Router /rewards [get]
func (h *RewardHandler) List(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	search := c.Query("search")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	rewards, err := h.rewardService.List(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rewards)
}

// Update @Summary Atualizar prêmio
// @Description Atualiza um prêmio existente
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Param reward body models.UpdateRewardRequest true "Dados para atualização"
// @Success 200 {object} models.RewardResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id} [put]
func (h *RewardHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	var req models.UpdateRewardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Dados inválidos",
			"message": err.Error(),
		})
		return
	}

	reward, err := h.rewardService.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reward)
}

// Delete @Summary Deletar prêmio
// @Description Remove um prêmio do sistema
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id} [delete]
func (h *RewardHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	if err := h.rewardService.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// AddBuyer @Summary Comprar números do prêmio
// @Description Compra uma quantidade específica de números para um usuário
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Param user_id path string true "ID do usuário"
// @Param request body models.BuyNumbersRequest true "Quantidade de números"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id}/buyers/{user_id} [post]
func (h *RewardHandler) AddBuyer(c *gin.Context) {
	rewardIDStr := c.Param("id")
	userIDStr := c.Param("user_id")

	rewardID, err := uuid.Parse(rewardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID do prêmio inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID do usuário inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	// Pegar a quantidade do body da requisição
	var req models.BuyNumbersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Quantidade é obrigatória",
			"message": err.Error(),
		})
		return
	}

	numbers, err := h.rewardService.BuyNumbers(rewardID, userID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Números comprados com sucesso",
		"numbers":  numbers,
		"quantity": req.Quantity,
	})
}

// RemoveBuyer @Summary Remover comprador do prêmio
// @Description Remove um usuário como comprador de um prêmio
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Param user_id path string true "ID do usuário"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id}/buyers/{user_id} [delete]
func (h *RewardHandler) RemoveBuyer(c *gin.Context) {
	rewardIDStr := c.Param("id")
	userIDStr := c.Param("user_id")

	rewardID, err := uuid.Parse(rewardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID do prêmio inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID do usuário inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	if err := h.rewardService.RemoveBuyer(rewardID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comprador removido com sucesso",
	})
}

// GetBuyers @Summary Listar compradores do prêmio
// @Description Lista todos os compradores de um prêmio específico
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Success 200 {array} models.UserResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id}/buyers [get]
func (h *RewardHandler) GetBuyers(c *gin.Context) {
	rewardIDStr := c.Param("id")

	rewardID, err := uuid.Parse(rewardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID do prêmio inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	buyers, err := h.rewardService.GetBuyers(rewardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, buyers)
}

// GetAvailableNumbers @Summary Buscar números disponíveis
// @Description Lista todos os números disponíveis para compra de um prêmio
// @Tags rewards
// @Accept json
// @Produce json
// @Param id path string true "ID do prêmio"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /rewards/{id}/available-numbers [get]
func (h *RewardHandler) GetAvailableNumbers(c *gin.Context) {
	rewardIDStr := c.Param("id")

	rewardID, err := uuid.Parse(rewardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "ID do prêmio inválido",
			"message": "Formato de ID inválido",
		})
		return
	}

	availableNumbers, err := h.rewardService.GetAvailableNumbers(rewardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Erro interno do servidor",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"available_numbers": availableNumbers,
		"total_available":   len(availableNumbers),
	})
}
