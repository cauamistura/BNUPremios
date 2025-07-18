# BNU Prêmios - Frontend

Frontend da aplicação BNU Prêmios desenvolvido em React com TypeScript, Vite e React Router.

## 🚀 Tecnologias

- **React 19.1.0** - Biblioteca para construção de interfaces
- **TypeScript 5.8.3** - Tipagem estática para JavaScript
- **Vite 6.3.5** - Build tool e dev server
- **React Router DOM 7.6.1** - Roteamento da aplicação
- **Styled Components 6.1.18** - Estilização com CSS-in-JS
- **React Slick 0.30.3** - Carrossel de imagens

## 📁 Estrutura do Projeto

```
src/
├── Components/           # Componentes reutilizáveis
│   ├── Header/          # Cabeçalho da aplicação
│   ├── ImageCarousel/   # Carrossel de imagens
│   ├── ProtectedRoute/  # Rota protegida por autenticação
│   ├── QuotaSelector/   # Seletor de quantidade
│   ├── RewardCard/      # Card de prêmio
│   ├── RewardCardList/  # Lista de cards de prêmios
│   ├── RewardForm/      # Formulário de criação/edição
│   ├── Toast/           # Componente de notificação
│   ├── ToastContainer/  # Container de notificações
│   └── TopBuyers/       # Lista de principais compradores
├── contexts/            # Contextos React
│   ├── AuthContext.tsx  # Contexto de autenticação
│   └── ToastContext.tsx # Contexto de notificações
├── hooks/               # Hooks customizados
│   ├── useAuth.ts       # Hook de autenticação
│   ├── useRedirectToLogin.ts # Hook de redirecionamento
│   └── useToast.ts      # Hook de notificações
├── Models/              # Interfaces TypeScript
│   ├── Buyers.tsx       # Interface de compradores
│   ├── Reaward.tsx      # Interface de prêmios
│   └── User.tsx         # Interface de usuários
├── Pages/               # Páginas da aplicação
│   ├── Auth/            # Páginas de autenticação
│   │   ├── Login/       # Página de login
│   │   └── Register/    # Página de registro
│   ├── Contacts/        # Página de contatos
│   ├── Home/            # Página inicial
│   ├── MyRewards/       # Página de meus prêmios
│   ├── Profile/         # Página de perfil
│   ├── RewardDetails/   # Página de detalhes do prêmio
│   └── RewardForm/      # Página de formulário
├── routes/              # Configuração de rotas
│   └── index.tsx        # Definição das rotas
├── services/            # Serviços de API
│   ├── apiUtils.ts      # Utilitários de requisição
│   ├── authService.ts   # Serviço de autenticação
│   ├── purchasesService.ts # Serviço de compras
│   ├── rewardsService.ts # Serviço de prêmios
│   └── usersService.ts  # Serviço de usuários
└── utils/               # Utilitários
    └── formatDate.ts    # Formatação de datas
```

## 🏗️ Arquitetura

### Contextos (Contexts)
- **AuthContext**: Gerencia estado de autenticação, login/logout e dados do usuário
- **ToastContext**: Sistema global de notificações com diferentes tipos (success, error, warning, info)

### Hooks Customizados
- **useAuth**: Acesso ao contexto de autenticação
- **useRedirectToLogin**: Redirecionamento para login quando não autenticado
- **useToast**: Sistema de notificações (deprecated em favor do ToastContext)

### Componentes Principais
- **ProtectedRoute**: Wrapper para rotas que requerem autenticação
- **RewardCard**: Card individual de prêmio com ações (editar, deletar, sortear)
- **RewardCardList**: Lista responsiva de cards de prêmios
- **Toast**: Notificação individual com animações
- **ToastContainer**: Container que gerencia múltiplas notificações

## 🔐 Sistema de Autenticação

### Funcionalidades
- Login/Logout com token JWT
- Proteção de rotas
- Redirecionamento automático
- Tratamento de token expirado
- Persistência de sessão no localStorage

### Tratamento de Token Expirado
```typescript
// Detecção automática de erros de token
if (errorMessage.includes('token is expired') || 
    errorMessage.includes('token has invalid claims') ||
    response.status === 401) {
  handleTokenExpired(); // Logout + redirecionamento
}
```

## 📱 Sistema de Notificações

### Toast Context
- Notificações globais com diferentes tipos
- Auto-dismiss configurável
- Animações de entrada/saída
- Posicionamento responsivo

### Tipos de Notificação
- **Success**: Ações bem-sucedidas
- **Error**: Erros e falhas
- **Warning**: Avisos importantes
- **Info**: Informações gerais

## 🎨 Estilização

### Abordagem
- CSS Modules para componentes específicos
- Styled Components para componentes reutilizáveis
- Design responsivo com breakpoints
- Gradientes e efeitos visuais modernos

### Características
- Design system consistente
- Cores temáticas (roxo como cor principal)
- Animações suaves
- Interface intuitiva

## 🔄 Gerenciamento de Estado

### Estados Locais
- Estados de loading para operações assíncronas
- Estados de formulários
- Estados de modais e confirmações

### Estados Globais
- Autenticação via Context API
- Notificações via Context API
- Dados de usuário persistidos

## 📡 Integração com API

### Serviços
- **rewardsService**: CRUD de prêmios, sorteios
- **authService**: Autenticação e registro
- **purchasesService**: Compras e números
- **usersService**: Dados de usuário

### Tratamento de Erros
- Captura de mensagens específicas da API
- Fallback para mensagens genéricas
- Logout automático em erros de autenticação
- Toasts para feedback do usuário

## 🚀 Scripts Disponíveis

```bash
# Desenvolvimento
npm run dev

# Build de produção
npm run build

# Preview do build
npm run preview

# Linting
npm run lint
```

## 📦 Dependências Principais

### Produção
- `react`: ^19.1.0
- `react-dom`: ^19.1.0
- `react-router-dom`: ^7.6.1
- `styled-components`: ^6.1.18
- `react-slick`: ^0.30.3

### Desenvolvimento
- `typescript`: ~5.8.3
- `vite`: ^6.3.5
- `eslint`: ^9.25.0
- `@vitejs/plugin-react`: ^4.4.1

## 🔧 Configurações

### TypeScript
- Configuração estrita para type safety
- Path mapping para imports limpos
- Configurações separadas para app e node

### Vite
- Plugin React para JSX
- Hot Module Replacement
- Build otimizado para produção

### ESLint
- Configuração para React e TypeScript
- Regras para hooks e refresh
- Integração com Vite

## 📱 Responsividade

### Breakpoints
- Mobile: < 768px
- Tablet: 768px - 1024px
- Desktop: > 1024px

### Adaptações
- Layout flexível para diferentes telas
- Componentes adaptáveis
- Navegação otimizada para mobile

## 🎯 Funcionalidades Implementadas

### Autenticação
- ✅ Login/Logout
- ✅ Registro de usuário
- ✅ Proteção de rotas
- ✅ Tratamento de token expirado

### Prêmios
- ✅ Listagem de prêmios
- ✅ Detalhes de prêmio
- ✅ Criação de prêmio
- ✅ Edição de prêmio
- ✅ Exclusão de prêmio
- ✅ Realização de sorteio
- ✅ Exibição do ganhador

### Compras
- ✅ Compra de números
- ✅ Seleção de quantidade
- ✅ Confirmação de compra
- ✅ Histórico de compras

### Notificações
- ✅ Sistema global de toasts
- ✅ Diferentes tipos de notificação
- ✅ Auto-dismiss configurável
- ✅ Animações suaves

## 🔍 Estrutura de Componentes

### Padrão de Nomenclatura
- Componentes em PascalCase
- Arquivos index.tsx para componentes principais
- CSS modules para estilos específicos

### Organização
- Componentes reutilizáveis em `/Components`
- Páginas em `/Pages`
- Lógica de negócio em `/services`
- Tipos em `/Models`

## 🚀 Deploy

### Build de Produção
```bash
npm run build
```

### Arquivos Gerados
- `dist/` - Arquivos otimizados para produção
- Assets com hash para cache busting
- Bundle otimizado e minificado

## 📝 Convenções

### Imports
```typescript
// Componentes
import ComponentName from '../../Components/ComponentName';

// Hooks
import { useAuth } from '../../hooks/useAuth';

// Serviços
import { rewardsService } from '../../services/rewardsService';

// Tipos
import type { Reward } from '../../Models/Reaward';
```

### Nomenclatura
- Interfaces: PascalCase (ex: `RewardDetails`)
- Funções: camelCase (ex: `handleSubmit`)
- Componentes: PascalCase (ex: `RewardCard`)
- Arquivos: kebab-case (ex: `reward-card.tsx`)

## 🔧 Desenvolvimento

### Setup Local
```bash
# Instalar dependências
npm install

# Iniciar servidor de desenvolvimento
npm run dev

# Acessar em http://localhost:5173
```

### Estrutura de Desenvolvimento
- Hot reload para desenvolvimento rápido
- TypeScript para type safety
- ESLint para qualidade de código
- Vite para build rápido

Este frontend foi desenvolvido com foco em performance, usabilidade e manutenibilidade, seguindo as melhores práticas do React e TypeScript.
