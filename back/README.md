# BNUPremios API

API REST para gerenciamento de usuários desenvolvida em Go com boas práticas de desenvolvimento.

## 🚀 Tecnologias

- **Go 1.21+**
- **Gin** - Framework web
- **PostgreSQL** - Banco de dados
- **Docker** - Containerização
- **Swagger** - Documentação da API
- **Golang Migrate** - Migrações do banco de dados

## 📋 Pré-requisitos

- Go 1.21 ou superior
- Docker e Docker Compose
- Git

## 🛠️ Instalação e Execução

### 1. Clone o repositório

```bash
git clone https://github.com/cauamistura/BNUPremios.git
cd BNUPremios/back
```

### 2. Configure as variáveis de ambiente

Copie o arquivo de exemplo e configure as variáveis:

```bash
cp env.example .env
```

Edite o arquivo `.env` com suas configurações:

```env
# Configurações do Banco de Dados
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=bnupremios
DB_SSLMODE=disable

# Configurações da API
API_PORT=8080
API_MODE=debug

# JWT Secret
JWT_SECRET=your-secret-key-here

# Configurações de Log
LOG_LEVEL=debug
```

### 3. Execute com Docker Compose

```bash
# Construir e executar os containers
docker-compose up --build

# Para executar em background
docker-compose up -d --build
```

### 4. Execute localmente (opcional)

Se preferir executar localmente:

```bash
# Instalar dependências
go mod download

# Executar migrações
go run cmd/main.go

# Ou executar a aplicação
go run cmd/main.go
```

## 📚 Documentação da API

A documentação Swagger está disponível em:

- **URL**: http://localhost:8080/swagger/index.html
- **Base URL**: http://localhost:8080/api/v1

## 🔧 Endpoints Disponíveis

### Autenticação
- `POST /api/v1/auth/register` - Registrar novo usuário
- `POST /api/v1/auth/login` - Login de usuário

### Usuários
- `GET /api/v1/users` - Listar usuários (com paginação)
- `POST /api/v1/users` - Criar usuário
- `GET /api/v1/users/{id}` - Buscar usuário por ID
- `PUT /api/v1/users/{id}` - Atualizar usuário
- `DELETE /api/v1/users/{id}` - Deletar usuário

### Health Check
- `GET /health` - Verificar status da API

## 📊 Estrutura do Projeto

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
│   │   └── user_handler.go    # Handlers HTTP
│   ├── middleware/
│   │   ├── cors.go           # Middleware CORS
│   │   └── logger.go         # Middleware de log
│   ├── models/
│   │   └── user.go           # Modelos de dados
│   ├── repository/
│   │   └── user_repository.go # Camada de acesso a dados
│   └── services/
│       └── user_service.go    # Lógica de negócio
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   └── 000001_create_users_table.down.sql
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── env.example
└── README.md
```

## 🗄️ Banco de Dados

### Tabela Users

| Campo      | Tipo      | Descrição                    |
|------------|-----------|------------------------------|
| id         | UUID      | Identificador único           |
| name       | VARCHAR   | Nome do usuário              |
| email      | VARCHAR   | Email (único)                |
| password   | VARCHAR   | Senha criptografada          |
| role       | VARCHAR   | Papel do usuário             |
| active     | BOOLEAN   | Status ativo/inativo         |
| created_at | TIMESTAMP | Data de criação              |
| updated_at | TIMESTAMP | Data de atualização          |

## 🔐 Segurança

- Senhas são criptografadas usando bcrypt
- Validação de dados de entrada
- Headers CORS configurados
- Logs de requisições

## 🧪 Testes

Para executar os testes (quando implementados):

```bash
go test ./...
```

## 📝 Exemplos de Uso

### Criar usuário

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "João Silva",
    "email": "joao@example.com",
    "password": "123456"
  }'
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@example.com",
    "password": "123456"
  }'
```

### Listar usuários

```bash
curl -X GET "http://localhost:8080/api/v1/users?page=1&limit=10"
```

## 🐛 Troubleshooting

### Problemas comuns

1. **Erro de conexão com banco de dados**
   - Verifique se o PostgreSQL está rodando
   - Confirme as configurações no arquivo `.env`

2. **Erro de migração**
   - Certifique-se de que o banco de dados está acessível
   - Verifique se as credenciais estão corretas

3. **Porta já em uso**
   - Altere a porta no arquivo `.env` ou `docker-compose.yml`

## 🤝 Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

**Cauã Mistura**
- GitHub: [@cauamistura](https://github.com/cauamistura)
- Projeto: [BNUPremios](https://github.com/cauamistura/BNUPremios) 