// @title           BNUPremios API
// @version         1.0
// @description     API para gerenciamento de prêmios e usuários do BNUPremios. Rotas públicas: listar prêmios, detalhes e compradores. Rotas protegidas: gerenciamento de usuários, criação/edição de prêmios, compra de números e listagem de compras.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Digite "Bearer" seguido de um espaço e o token JWT.

package main

import (
	"log"
	"os"

	"github.com/cauamistura/BNUPremios/internal/config"
	"github.com/cauamistura/BNUPremios/internal/database"
	"github.com/cauamistura/BNUPremios/internal/handlers"
	"github.com/cauamistura/BNUPremios/internal/repository"
	"github.com/cauamistura/BNUPremios/internal/routes"
	"github.com/cauamistura/BNUPremios/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {

	// Carregar configurações
	cfg := config.Load()

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Executar migrações
	if err := database.RunMigrations(cfg.Database); err != nil {
		log.Fatal("Erro ao executar migrações:", err)
	}

	// Configurar repositórios
	userRepo := repository.NewUserRepository(db)

	// Configurar serviços
	userService := services.NewUserService(userRepo, cfg.JWT.Secret)

	// Configurar handlers
	userHandler := handlers.NewUserHandler(userService)
	rewardHandler := handlers.NewRewardHandler(services.NewRewardService(repository.NewRewardRepository(db)))

	// Configurar Gin
	if cfg.API.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Configurar rotas
	routes.SetupRoutes(router, userHandler, rewardHandler, cfg.JWT.Secret)

	// Iniciar servidor
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciado na porta %s", port)
	log.Printf("Acesse http://localhost:%s/swagger/index.html para a documentação", port)
	log.Fatal(router.Run(":" + port))
}
