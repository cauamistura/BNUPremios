.PHONY: help build run test clean docker-build docker-run docker-stop swagger

# Variáveis
APP_NAME=bnupremios
MAIN_FILE=cmd/main.go
DOCKER_IMAGE=bnupremios-api

# Comando padrão
help: ## Mostra esta ajuda
	@echo "Comandos disponíveis:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAFEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Desenvolvimento local
build: ## Compila a aplicação
	go build -o bin/$(APP_NAME) $(MAIN_FILE)

run: ## Executa a aplicação localmente
	go run $(MAIN_FILE)

test: ## Executa os testes
	go test ./...

clean: ## Remove arquivos temporários
	go clean
	rm -rf bin/
	rm -rf docs/

# Docker
docker-build: ## Constrói a imagem Docker
	docker build -t $(DOCKER_IMAGE) .

docker-run: ## Executa com Docker Compose
	docker-compose up --build

docker-stop: ## Para os containers Docker
	docker-compose down

# Swagger
swagger: ## Gera a documentação Swagger
	swag init -g $(MAIN_FILE)

# Migrações
migrate-up: ## Executa migrações para cima
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/bnupremios?sslmode=disable" up

migrate-down: ## Executa migrações para baixo
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/bnupremios?sslmode=disable" down

# Dependências
deps: ## Instala dependências
	go mod download
	go mod tidy

# Desenvolvimento
dev: deps swagger run ## Executa em modo desenvolvimento

# Produção
prod: build ## Prepara para produção
	@echo "Aplicação compilada em bin/$(APP_NAME)"

# Instalação do Swagger CLI (se necessário)
install-swagger: ## Instala o Swagger CLI
	go install github.com/swaggo/swag/cmd/swag@latest

# Verificação de código
lint: ## Executa linter (se configurado)
	@echo "Linting não configurado ainda"

# Setup inicial
setup: install-swagger deps swagger ## Configuração inicial do projeto
	@echo "Projeto configurado com sucesso!" 