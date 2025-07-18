# 🎰 BNUPremios - Plataforma de Sorteios

Uma plataforma completa para criação, gerenciamento e participação em sorteios online. Desenvolvida com arquitetura moderna, permite que usuários criem prêmios, comprem números e realizem sorteios de forma segura e transparente.

## 🎯 Sobre o Projeto

O BNUPremios é uma solução completa para sorteios digitais que oferece:

- **Criação de Prêmios**: Usuários podem criar prêmios com descrição, valor e imagens
- **Compra de Números**: Sistema de compra de números para participação nos sorteios
- **Sorteios Automáticos**: Realização de sorteios com seleção aleatória de ganhadores
- **Gestão Completa**: Dashboard para gerenciar prêmios criados e compras realizadas
- **Interface Moderna**: Design responsivo e intuitivo para melhor experiência do usuário

## 🏗️ Arquitetura

O projeto utiliza uma arquitetura full-stack moderna:

```
┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │
│   React + TS    │◄──►│   Go + Gin      │
│   Vite          │    │   PostgreSQL    │
└─────────────────┘    └─────────────────┘
```

### Stack Tecnológica

#### Frontend
- **React 19.1.0** - Interface de usuário moderna
- **TypeScript 5.8.3** - Tipagem estática para maior segurança
- **Vite 6.3.5** - Build tool rápido e eficiente
- **React Router DOM 7.6.1** - Navegação entre páginas
- **Styled Components 6.1.18** - Estilização com CSS-in-JS
- **React Slick 0.30.3** - Carrossel de imagens para prêmios

#### Backend
- **Go 1.24.5** - Linguagem de programação performática
- **Gin** - Framework web para APIs RESTful
- **PostgreSQL** - Banco de dados relacional robusto
- **JWT** - Autenticação segura com tokens
- **Golang Migrate** - Gerenciamento de migrações de banco
- **Docker** - Containerização para deploy

## 🚀 Funcionalidades Principais

### 👤 Sistema de Usuários
- Registro e login seguro
- Perfil de usuário personalizado
- Histórico de compras e prêmios criados

### 🎁 Gestão de Prêmios
- Criação de prêmios com imagens e descrições
- Definição de valores e quantidades
- Edição e exclusão de prêmios próprios
- Listagem pública de prêmios disponíveis

### 🎫 Sistema de Compra
- Seleção de quantidade de números
- Confirmação de compra
- Atribuição automática de números
- Histórico de compras por usuário

### 🎲 Sorteios
- Realização de sorteios automáticos
- Seleção aleatória de ganhadores
- Exibição do resultado do sorteio
- Transparência no processo

### 📊 Dashboard
- Visão geral dos prêmios criados
- Estatísticas de compras
- Lista de principais compradores
- Gestão completa de sorteios

## 🛠️ Como Executar

### Pré-requisitos
- Node.js 18+ e npm
- Go 1.24.5+
- PostgreSQL 15+
- Docker (opcional)

### Configuração Rápida

1. **Clone o repositório**
   ```bash
   git clone <repository-url>
   cd BNUPremios
   ```

2. **Configure o Backend**
   ```bash
   cd back
   cp env.copy.example .env
   # Edite o arquivo .env com suas configurações
   make setup
   make run
   ```

3. **Configure o Frontend**
   ```bash
   cd front
   npm install
   npm run dev
   ```

4. **Acesse a aplicação**
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080
   - Documentação: http://localhost:8080/swagger/index.html

### Execução com Docker
```bash
# Backend
cd back
make docker-run

# Frontend
cd front
npm run build
# Servir os arquivos estáticos
```

## 📱 Principais Páginas

### Públicas
- **Home** - Listagem de prêmios disponíveis
- **Detalhes do Prêmio** - Informações completas e compra de números
- **Login/Registro** - Autenticação de usuários

### Protegidas
- **Meus Prêmios** - Gestão de prêmios criados
- **Perfil** - Dados pessoais e histórico
- **Criar Prêmio** - Formulário de criação
- **Contatos** - Informações de suporte

## 🔐 Segurança

- Autenticação JWT com tokens seguros
- Senhas criptografadas com bcrypt
- Validação de entrada em todos os endpoints
- Middleware CORS configurado
- Proteção de rotas no frontend

## 📊 Banco de Dados

### Principais Tabelas
- **users** - Dados dos usuários
- **rewards** - Informações dos prêmios
- **reward_buyers** - Relacionamento prêmio-comprador

### Migrações
```bash
cd back
make migrate-up    # Executar migrações
make migrate-down  # Reverter migrações
```

## 🎨 Interface

- Design responsivo para mobile, tablet e desktop
- Tema em tons de roxo com gradientes modernos
- Animações suaves e feedback visual
- Sistema de notificações toast
- Carrossel de imagens para prêmios

## 🔧 Desenvolvimento

### Comandos Úteis

```bash
# Backend
make help          # Ver todos os comandos
make dev           # Desenvolvimento completo
make test          # Executar testes
make swagger       # Gerar documentação

# Frontend
npm run dev        # Servidor de desenvolvimento
npm run build      # Build de produção
npm run preview    # Preview do build
npm run lint       # Verificar código
```

### Estrutura de Código
- **Arquitetura em camadas** no backend (Handlers → Services → Repositories)
- **Componentes reutilizáveis** no frontend
- **Context API** para gerenciamento de estado global
- **Hooks customizados** para lógica reutilizável

## 📈 Próximas Funcionalidades

- [ ] Sistema de pagamentos integrado
- [ ] Notificações em tempo real
- [ ] Relatórios e analytics
- [ ] Sistema de convites
- [ ] Integração com redes sociais
- [ ] App mobile nativo

---

**BNUPremios** - Transformando sorteios em experiências digitais seguras e transparentes! 🎰✨