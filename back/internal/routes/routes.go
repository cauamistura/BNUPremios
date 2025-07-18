package routes

import (
	"github.com/cauamistura/BNUPremios/docs"
	"github.com/cauamistura/BNUPremios/internal/handlers"
	"github.com/cauamistura/BNUPremios/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(router *gin.Engine, userHandler *handlers.UserHandler, rewardHandler *handlers.RewardHandler, jwtSecret string) {
	// Middleware global
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "BNUPremios API está funcionando",
		})
	})

	// Swagger
	docs.SwaggerInfo.Title = "BNUPremios API"
	docs.SwaggerInfo.Description = "API para gerenciamento de usuários do BNUPremios"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Grupo de rotas da API
	api := router.Group("/api/v1")
	{
		// Rotas de usuários
		users := api.Group("/users")
		users.Use(middleware.AuthMiddleware(jwtSecret))
		{
			users.GET("/", userHandler.List)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}

		// Rotas de compras (protegidas por autenticação)
		purchases := api.Group("/purchases")
		purchases.Use(middleware.AuthMiddleware(jwtSecret))
		{
			purchases.GET("/user/:user_id", rewardHandler.GetUserPurchases)
		}

		// Rotas de autenticação
		auth := api.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
			auth.POST("/register", userHandler.Register)
		}

		// Rotas de prêmios
		rewards := api.Group("/rewards")
		{
			// Rotas públicas (sem autenticação)
			rewards.GET("/", rewardHandler.List)
			rewards.GET("/:id", rewardHandler.GetByID)
			rewards.GET("/:id/details", rewardHandler.GetDetailsByID)
			rewards.GET("/:id/buyers", rewardHandler.GetBuyers)

			// Rotas protegidas (com autenticação)
			protectedRewards := rewards.Group("/")
			protectedRewards.Use(middleware.AuthMiddleware(jwtSecret))
			{
				protectedRewards.POST("/", rewardHandler.Create)
				protectedRewards.GET("/mine", rewardHandler.ListMyRewards)
				protectedRewards.PUT("/:id", rewardHandler.Update)
				protectedRewards.DELETE("/:id", rewardHandler.Delete)

				// Rotas de compradores protegidas
				protectedRewards.POST("/:id/buyers/:user_id", rewardHandler.AddBuyer)
				protectedRewards.DELETE("/:id/buyers/:user_id", rewardHandler.RemoveBuyer)
				protectedRewards.GET("/:id/buyers/:user_id/numbers", rewardHandler.GetUserNumbers)
				protectedRewards.POST("/:id/draw", rewardHandler.Draw)
			}
		}
	}
}
