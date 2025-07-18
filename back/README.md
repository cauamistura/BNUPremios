# BNUPremios API

API RESTful desenvolvida em Go para gerenciamento de prêmios e usuários do sistema BNUPremios.

## 🏗️ Arquitetura

A API segue uma arquitetura em camadas (Layered Architecture) com separação clara de responsabilidades:

```
┌─────────────────┐
│   Controllers   │  ← Handlers (Gin)
├─────────────────┤
│    Services     │  ← Lógica de Negócio
├─────────────────┤
│  Repositories   │  ← Acesso a Dados
├─────────────────┤
│   Database      │  ← PostgreSQL
└─────────────────┘
```

### Estrutura do Projeto

```
back/
├── cmd/
│   └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── config/
│   │   └── config.go          # Configurações da aplicação
│   ├── database/
│   │   └── connection.go      # Conexão com banco de dados
│   ├── handlers/
│   │   ├── user_handler.go    # Controllers de usuários
│   │   └── reward_handler.go  # Controllers de prêmios
│   ├── middleware/
│   │   ├── auth.go           # Middleware de autenticação JWT
│   │   ├── cors.go           # Middleware CORS
│   │   └── logger.go         # Middleware de logging
│   ├── models/
│   │   ├── user.go           # Modelos de usuário
│   │   └── reward.go         # Modelos de prêmio
│   ├── repository/
│   │   ├── user_repository.go # Repositório de usuários
│   │   └── reward_repository.go # Repositório de prêmios
│   ├── routes/
│   │   └── routes.go         # Definição de rotas
│   └── services/
│       ├── user_service.go   # Serviços de usuário
│       └── reward_service.go # Serviços de prêmio
├── migrations/               # Migrações do banco de dados
├── Dockerfile               # Configuração Docker
├── docker-compose.yml       # Orquestração de containers
├── go.mod                   # Dependências Go
└── Makefile                 # Comandos de automação
```

## 🛠️ Tecnologias Utilizadas

- **Go 1.24.5** - Linguagem principal
- **Gin** - Framework web para roteamento e middleware
- **PostgreSQL** - Banco de dados relacional
- **JWT** - Autenticação e autorização
- **Swagger** - Documentação da API
- **Docker** - Containerização
- **Golang Migrate** - Migrações de banco de dados
- **UUID** - Identificadores únicos

## 🚀 Como Executar

### Pré-requisitos

- Go 1.24.5 ou superior
- PostgreSQL 15 ou superior
- Docker e Docker Compose (opcional)

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto:

```env
# Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=bnupremios
DB_SSLMODE=disable

# API
API_PORT=8080
API_MODE=debug

# JWT
JWT_SECRET=your-secret-key-here
```

### Execução Local

1. **Instalar dependências:**
   ```bash
   go mod download
   go mod tidy
   ```

2. **Executar migrações:**
   ```bash
   make migrate-up
   ```

3. **Executar a aplicação:**
   ```bash
   make run
   ```

### Execução com Docker

```bash
# Construir e executar com Docker Compose
make docker-run

# Ou manualmente:
docker-compose up --build
```

## 📚 Endpoints da API

### Autenticação
- `POST /api/v1/auth/login` - Login de usuário
- `POST /api/v1/auth/register` - Registro de usuário

### Usuários (Protegido)
- `GET /api/v1/users/` - Listar usuários
- `GET /api/v1/users/:id` - Obter usuário por ID
- `PUT /api/v1/users/:id` - Atualizar usuário
- `DELETE /api/v1/users/:id` - Deletar usuário

### Prêmios
#### Públicos
- `GET /api/v1/rewards/` - Listar prêmios
- `GET /api/v1/rewards/:id` - Obter prêmio por ID
- `GET /api/v1/rewards/:id/details` - Obter detalhes do prêmio
- `GET /api/v1/rewards/:id/buyers` - Listar compradores

#### Protegidos
- `POST /api/v1/rewards/` - Criar prêmio
- `GET /api/v1/rewards/mine` - Listar meus prêmios
- `PUT /api/v1/rewards/:id` - Atualizar prêmio
- `DELETE /api/v1/rewards/:id` - Deletar prêmio
- `POST /api/v1/rewards/:id/buyers/:user_id` - Adicionar comprador
- `DELETE /api/v1/rewards/:id/buyers/:user_id` - Remover comprador
- `GET /api/v1/rewards/:id/buyers/:user_id/numbers` - Obter números do usuário
- `POST /api/v1/rewards/:id/draw` - Realizar sorteio

### Compras (Protegido)
- `GET /api/v1/purchases/user/:user_id` - Listar compras do usuário

### Utilitários
- `GET /health` - Health check
- `GET /swagger/*` - Documentação Swagger

## 🔐 Autenticação

A API utiliza JWT (JSON Web Tokens) para autenticação. Para acessar endpoints protegidos:

1. Faça login via `POST /api/v1/auth/login`
2. Use o token retornado no header: `Authorization: Bearer <token>`

## 📊 Banco de Dados

### Tabelas Principais

- **users** - Usuários do sistema
- **rewards** - Prêmios disponíveis
- **reward_buyers** - Relacionamento entre prêmios e compradores

### Migrações

As migrações são gerenciadas com `golang-migrate`:

```bash
# Executar migrações
make migrate-up

# Reverter migrações
make migrate-down
```

## 🐳 Docker

### Construir Imagem
```bash
make docker-build
```

### Executar com Docker Compose
```bash
make docker-run
```

### Parar Containers
```bash
make docker-stop
```

## 📖 Documentação

A documentação da API está disponível via Swagger em:
```
http://localhost:8080/swagger/index.html
```

Para gerar a documentação:
```bash
make swagger
```

## 🧪 Desenvolvimento

### Comandos Úteis

```bash
# Ajuda
make help

# Desenvolvimento completo
make dev

# Testes
make test

# Limpeza
make clean

# Setup inicial
make setup
```

### Estrutura de Código

- **Handlers**: Responsáveis por receber requisições HTTP e retornar respostas
- **Services**: Contêm a lógica de negócio da aplicação
- **Repositories**: Gerenciam o acesso aos dados no banco
- **Models**: Definem as estruturas de dados
- **Middleware**: Interceptam requisições para autenticação, CORS, logging, etc.

## 🔧 Configuração

A aplicação utiliza configuração baseada em variáveis de ambiente com fallbacks para valores padrão. As configurações são carregadas no pacote `config`.

### Modos de Execução

- **debug**: Modo de desenvolvimento com logs detalhados
- **release**: Modo de produção com logs mínimos

## 📝 Logs

A aplicação utiliza middleware de logging que registra:
- Método HTTP
- URL da requisição
- Status da resposta
- Tempo de processamento
- Tamanho da resposta

## 🔒 Segurança

- Autenticação JWT
- Senhas criptografadas com bcrypt
- Middleware CORS configurado
- Validação de entrada com tags de binding
- UUIDs para identificadores

## 🚀 Deploy

A aplicação está preparada para deploy em containers Docker com:
- Multi-stage build para otimização
- Imagem Alpine para tamanho reduzido
- Configuração via variáveis de ambiente
- Health check endpoint

## 📄 Licença

Este projeto está sob a licença Apache 2.0.
