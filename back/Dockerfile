FROM golang:1.24-alpine AS builder

WORKDIR /app

# Instalar dependências necessárias
RUN apk add --no-cache git

# Copiar go mod e sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Instalar swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Gerar documentação Swagger
RUN swag init -g cmd/main.go

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Imagem final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiar o binário da aplicação
COPY --from=builder /app/main .

# Copiar a documentação Swagger
COPY --from=builder /app/docs ./docs

# Expor porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"] 