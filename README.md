# RYV API - Sistema de Blog com Painel Administrativo

API REST para gerenciamento de blog com sistema de autenticação JWT e painel administrativo seguro.

## 🚀 Funcionalidades

- **Blog Público**: Artigos, categorias e recomendações diárias
- **Sistema de Autenticação**: JWT com middleware de segurança
- **Painel Administrativo**: Gerenciamento completo de conteúdo
- **WhatsApp Integration**: Sistema de contatos e estatísticas
- **Segurança**: Middlewares de autenticação, rate limiting e CORS

## 🛠️ Tecnologias

- **Go 1.23+** - Linguagem principal
- **Gin** - Framework web
- **GORM** - ORM para banco de dados
- **SQLite** - Banco de dados
- **JWT** - Autenticação
- **Docker** - Containerização

## 📋 Pré-requisitos

- Go 1.23 ou superior
- Docker e Docker Compose (opcional)

## 🚀 Instalação

### Opção 1: Execução Local

1. **Clone o repositório**

```bash
git clone <repository-url>
cd ryv-api
```

2. **Instale as dependências**

```bash
go mod tidy
```

3. **Configure as variáveis de ambiente**

```bash
cp env.example .env
# Edite o arquivo .env com suas configurações
```

4. **Execute a aplicação**

```bash
go run main.go
```

### Opção 2: Docker Compose

1. **Clone e configure**

```bash
git clone <repository-url>
cd ryv-api
```

2. **Execute com Docker**

```bash
docker-compose up --build
```

## 🔧 Configuração Inicial

### 1. Criar Primeiro Administrador

Execute o script para criar o primeiro admin:

```bash
go run scripts/create-admin.go
```

Ou use a API diretamente:

```bash
curl -X POST http://localhost:3001/api/auth/create-admin \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin",
    "email": "admin@example.com",
    "password": "senha123"
  }'
```

### 2. Fazer Login

```bash
curl -X POST http://localhost:3001/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "senha123"
  }'
```

## 📚 Endpoints da API

### 🔓 Rotas Públicas

#### Artigos

- `GET /api/articles` - Listar artigos publicados
- `GET /api/articles/categories` - Listar categorias
- `GET /api/articles/:slug` - Buscar artigo por slug
- `GET /api/articles/daily-recommendation` - Recomendação diária

#### WhatsApp

- `POST /api/whatsapp/contact` - Registrar contato

### 🔐 Rotas de Autenticação

- `POST /api/auth/login` - Fazer login
- `POST /api/auth/register` - Registrar usuário
- `POST /api/auth/create-admin` - Criar admin (setup inicial)

### 🛡️ Rotas Protegidas (Admin)

**Header necessário**: `Authorization: Bearer <token>`

#### Perfil

- `GET /api/admin/profile` - Perfil do usuário

#### Gerenciamento de Artigos

- `POST /api/admin/articles` - Criar artigo
- `PUT /api/admin/articles/:id` - Atualizar artigo
- `DELETE /api/admin/articles/:id` - Deletar artigo

#### WhatsApp (Admin)

- `GET /api/admin/whatsapp/contacts` - Listar contatos
- `GET /api/admin/whatsapp/stats` - Estatísticas

## 🔒 Segurança

### Middlewares Implementados

1. **AuthMiddleware**: Validação de JWT
2. **AdminMiddleware**: Verificação de permissões de admin
3. **RateLimitMiddleware**: Proteção contra ataques de força bruta
4. **CORS**: Configuração de origens permitidas

### Configurações de Segurança

- Tokens JWT com expiração de 24 horas
- Senhas hasheadas com bcrypt
- Headers de segurança configurados
- Validação de entrada em todos os endpoints

## 📊 Estrutura do Projeto

```
ryv-api/
├── database/          # Configuração do banco de dados
├── handlers/          # Handlers da API
├── middleware/        # Middlewares de segurança
├── models/           # Modelos de dados
├── scripts/          # Scripts utilitários
├── scraper/          # Sistema de scraping
├── seed/             # Dados iniciais
├── main.go           # Arquivo principal
├── docker-compose.yml # Configuração Docker
└── README.md         # Documentação
```

## 🐳 Docker

### Build da Imagem

```bash
docker build -t ryv-api .
```

### Execução com Docker Compose

```bash
docker-compose up -d
```

### Logs

```bash
docker-compose logs -f ryv-api
```

## 🔧 Desenvolvimento

### Executar em modo desenvolvimento

```bash
export GIN_MODE=debug
go run main.go
```

### Executar testes

```bash
go test ./...
```

### Formatar código

```bash
go fmt ./...
```

## 📝 Exemplos de Uso

### Criar um Artigo (Admin)

```bash
curl -X POST http://localhost:3001/api/admin/articles \
  -H "Authorization: Bearer <seu-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Como Cuidar da Saúde Ocular",
    "slug": "como-cuidar-saude-ocular",
    "content": "Conteúdo do artigo...",
    "excerpt": "Resumo do artigo",
    "category": "Optometria",
    "tags": "saúde,olhos,prevenção",
    "author": "Dr. João Silva",
    "is_published": true
  }'
```

### Registrar Contato WhatsApp

```bash
curl -X POST http://localhost:3001/api/whatsapp/contact \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Maria Silva",
    "phone": "5511999999999",
    "message": "Gostaria de agendar uma consulta",
    "source": "artigo-saude-ocular"
  }'
```

## 🚨 Troubleshooting

### Erro de Conexão com Banco

- Verifique se o arquivo `ryv_blog.db` existe
- Execute `go run main.go` para criar o banco automaticamente

### Erro de Autenticação

- Verifique se o `JWT_SECRET` está configurado
- Confirme se o token não expirou (24h)

### Erro de CORS

- Verifique as configurações de `ALLOWED_ORIGINS`
- Confirme se o frontend está na origem permitida

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📞 Suporte

Para suporte, envie um email para [seu-email@exemplo.com] ou abra uma issue no GitHub.
