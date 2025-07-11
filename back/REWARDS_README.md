# CRUD de Prêmios (Rewards) - BNUPremios API

Este documento descreve as funcionalidades implementadas para o CRUD de prêmios no sistema BNUPremios.

## Modelos Implementados

### Reward (Prêmio Básico)
```go
type Reward struct {
    ID          uuid.UUID `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Image       string    `json:"image"`
    DrawDate    time.Time `json:"draw_date"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### RewardDetails (Detalhes do Prêmio)
```go
type RewardDetails struct {
    Reward
    Images   []string `json:"images"`
    Price    float64  `json:"price"`
    MinQuota int      `json:"min_quota"`
    Buyers   []User   `json:"buyers"`
}
```

## Endpoints da API

### 1. Criar Prêmio
- **POST** `/api/v1/rewards`
- **Body:**
```json
{
    "name": "iPhone 15 Pro",
    "description": "Smartphone Apple iPhone 15 Pro 128GB",
    "image": "https://example.com/iphone.jpg",
    "draw_date": "2024-12-31T23:59:59Z",
    "images": ["https://example.com/iphone1.jpg", "https://example.com/iphone2.jpg"],
    "price": 8999.99,
    "min_quota": 100
}
```

### 2. Listar Prêmios
- **GET** `/api/v1/rewards`
- **Query Parameters:**
  - `page` (opcional): Número da página (padrão: 1)
  - `limit` (opcional): Itens por página (padrão: 10, máximo: 100)
  - `search` (opcional): Termo de busca

### 3. Buscar Prêmio por ID
- **GET** `/api/v1/rewards/{id}`

### 4. Buscar Detalhes do Prêmio
- **GET** `/api/v1/rewards/{id}/details`
- Retorna o prêmio com todas as informações detalhadas (imagens, preço, quota mínima e compradores)

### 5. Atualizar Prêmio
- **PUT** `/api/v1/rewards/{id}`
- **Body:** Campos opcionais para atualização
```json
{
    "name": "iPhone 15 Pro Max",
    "completed": true,
    "price": 9999.99
}
```

### 6. Deletar Prêmio
- **DELETE** `/api/v1/rewards/{id}`

### 7. Gerenciar Compradores

#### Adicionar Comprador
- **POST** `/api/v1/rewards/{id}/buyers/{user_id}`

#### Remover Comprador
- **DELETE** `/api/v1/rewards/{id}/buyers/{user_id}`

#### Listar Compradores
- **GET** `/api/v1/rewards/{id}/buyers`

## Estrutura do Banco de Dados

### Tabela `rewards`
- `id` (UUID, PK)
- `name` (VARCHAR(255), NOT NULL)
- `description` (TEXT)
- `image` (VARCHAR(500))
- `draw_date` (TIMESTAMP, NOT NULL)
- `completed` (BOOLEAN, DEFAULT FALSE)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

### Tabela `reward_details`
- `reward_id` (UUID, PK, FK para rewards)
- `price` (DECIMAL(10,2), DEFAULT 0.00)
- `min_quota` (INTEGER, DEFAULT 1)
- `created_at` (TIMESTAMP)
- `updated_at` (TIMESTAMP)

### Tabela `reward_images`
- `id` (UUID, PK)
- `reward_id` (UUID, FK para rewards)
- `image_url` (VARCHAR(500), NOT NULL)
- `created_at` (TIMESTAMP)

### Tabela `reward_buyers`
- `reward_id` (UUID, FK para rewards)
- `user_id` (UUID, FK para users)
- `created_at` (TIMESTAMP)
- **Primary Key:** (reward_id, user_id)

## Funcionalidades Implementadas

### ✅ CRUD Completo
- ✅ Criar prêmio
- ✅ Listar prêmios com paginação
- ✅ Buscar prêmio por ID
- ✅ Atualizar prêmio
- ✅ Deletar prêmio

### ✅ Detalhes do Prêmio
- ✅ Buscar detalhes completos (imagens, preço, quota mínima)
- ✅ Gerenciar imagens adicionais
- ✅ Definir preço e quota mínima

### ✅ Gerenciamento de Compradores
- ✅ Adicionar comprador ao prêmio
- ✅ Remover comprador do prêmio
- ✅ Listar todos os compradores de um prêmio

### ✅ Funcionalidades Adicionais
- ✅ Busca por nome e descrição
- ✅ Paginação com metadados
- ✅ Validação de dados
- ✅ Tratamento de erros
- ✅ Documentação Swagger completa

## Arquivos Criados/Modificados

### Novos Arquivos
- `internal/models/reward.go` - Modelos de dados
- `internal/repository/reward_repository.go` - Camada de acesso a dados
- `internal/services/reward_service.go` - Lógica de negócio
- `internal/handlers/reward_handler.go` - Controladores HTTP
- `migrations/000002_create_rewards_table.up.sql` - Migração de criação
- `migrations/000002_create_rewards_table.down.sql` - Migração de rollback

### Arquivos Modificados
- `internal/routes/routes.go` - Adicionadas rotas de prêmios
- `cmd/main.go` - Configuração do RewardHandler

## Como Usar

1. **Executar migrações:**
```bash
make migrate-up
```

2. **Iniciar o servidor:**
```bash
make run
```

3. **Acessar documentação:**
```
http://localhost:8080/swagger/index.html
```

## Exemplo de Uso

### Criar um Prêmio
```bash
curl -X POST http://localhost:8080/api/v1/rewards \
  -H "Content-Type: application/json" \
  -d '{
    "name": "PlayStation 5",
    "description": "Console Sony PlayStation 5",
    "image": "https://example.com/ps5.jpg",
    "draw_date": "2024-12-31T23:59:59Z",
    "images": ["https://example.com/ps5_1.jpg", "https://example.com/ps5_2.jpg"],
    "price": 3999.99,
    "min_quota": 50
  }'
```

### Listar Prêmios
```bash
curl "http://localhost:8080/api/v1/rewards?page=1&limit=10&search=PlayStation"
```

### Adicionar Comprador
```bash
curl -X POST http://localhost:8080/api/v1/rewards/{reward_id}/buyers/{user_id}
```

## Relacionamento com Usuários

O sistema já possui um CRUD completo de usuários (`User`), e os compradores (`Buyer`) são na verdade os usuários do sistema. Quando um usuário se torna comprador de um prêmio, ele é adicionado à tabela `reward_buyers`, criando o relacionamento entre prêmios e usuários.

Isso permite:
- Rastrear quais usuários compraram cada prêmio
- Listar todos os compradores de um prêmio específico
- Gerenciar a lista de compradores dinamicamente 