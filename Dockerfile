# Multi-stage build para otimizar tamanho da imagem
FROM golang:1.23-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git ca-certificates tzdata

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências primeiro (para cache)
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o ryv-api main.go

# Imagem final
FROM alpine:latest

# Instalar dependências necessárias para SQLite
RUN apk --no-cache add ca-certificates tzdata sqlite

# Criar usuário não-root para segurança
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Definir diretório de trabalho
WORKDIR /app

# Copiar apenas o binário compilado
COPY --from=builder /app/ryv-api .

# Copiar arquivos necessários
COPY --from=builder /app/database ./database
COPY --from=builder /app/models ./models
COPY --from=builder /app/middleware ./middleware
COPY --from=builder /app/handlers ./handlers

# Definir permissões
RUN chown -R appuser:appgroup /app

# Mudar para usuário não-root
USER appuser

# Expor porta
EXPOSE 3001

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3001/health || exit 1

# Comando para executar a aplicação
CMD ["./ryv-api"] 