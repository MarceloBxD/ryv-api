.PHONY: build run test clean docker-build docker-run create-admin help

# Variáveis
BINARY_NAME=ryv-api
MAIN_FILE=main.go
DOCKER_IMAGE=ryv-api

# Comandos principais
build:
	@echo "🔨 Compilando aplicação..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Build concluído!"

run:
	@echo "🚀 Executando aplicação..."
	go run $(MAIN_FILE)

dev:
	@echo "🔧 Executando em modo desenvolvimento..."
	GIN_MODE=debug go run $(MAIN_FILE)

admin:
	@echo "🌐 Iniciando painel administrativo..."
	cd admin && go run server.go

test:
	@echo "🧪 Executando testes..."
	go test ./...

clean:
	@echo "🧹 Limpando arquivos..."
	rm -f $(BINARY_NAME)
	go clean

# Docker
docker-build:
	@echo "🐳 Construindo imagem Docker..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "🐳 Executando com Docker Compose..."
	docker-compose up --build

docker-stop:
	@echo "🛑 Parando containers..."
	docker-compose down

docker-logs:
	@echo "📋 Exibindo logs..."
	docker-compose logs -f

# Administração
create-admin:
	@echo "👤 Criando administrador..."
	go run scripts/create-admin.go

# Dependências
deps:
	@echo "📦 Baixando dependências..."
	go mod tidy
	go mod download

# Formatação e linting
fmt:
	@echo "🎨 Formatando código..."
	go fmt ./...

lint:
	@echo "🔍 Verificando código..."
	golangci-lint run

# Banco de dados
db-reset:
	@echo "🗄️ Resetando banco de dados..."
	rm -f ryv_blog.db
	go run $(MAIN_FILE)

# Ajuda
help:
	@echo "📚 Comandos disponíveis:"
	@echo ""
	@echo "🔨 Build e Execução:"
	@echo "  make build        - Compilar a aplicação"
	@echo "  make run          - Executar a aplicação"
	@echo "  make dev          - Executar em modo desenvolvimento"
	@echo "  make clean        - Limpar arquivos compilados"
	@echo ""
	@echo "🐳 Docker:"
	@echo "  make docker-build - Construir imagem Docker"
	@echo "  make docker-run   - Executar com Docker Compose"
	@echo "  make docker-stop  - Parar containers"
	@echo "  make docker-logs  - Exibir logs dos containers"
	@echo ""
	@echo "🧪 Testes e Qualidade:"
	@echo "  make test         - Executar testes"
	@echo "  make fmt          - Formatando código"
	@echo "  make lint         - Verificar código"
	@echo ""
	@echo "👤 Administração:"
	@echo "  make create-admin - Criar primeiro administrador"
	@echo "  make db-reset     - Resetar banco de dados"
	@echo ""
	@echo "📦 Dependências:"
	@echo "  make deps         - Baixar e organizar dependências"
	@echo ""
	@echo "❓ Ajuda:"
	@echo "  make help         - Exibir esta ajuda"

# Comando padrão
.DEFAULT_GOAL := help 