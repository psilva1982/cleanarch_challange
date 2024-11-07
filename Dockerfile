# Etapa de build
FROM golang:1.23.2-bullseye AS builder

# Instala o Git e o Wire
RUN apt-get update && apt-get install -y git && \
    go install github.com/google/wire/cmd/wire@latest && \
    rm -rf /var/lib/apt/lists/*

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos do projeto e o arquivo .env para o contêiner
COPY . .
COPY cmd/ordersystem/.env /app/cmd/ordersystem/.env

# Define o diretório para build
WORKDIR /app/cmd/ordersystem

# Baixa as dependências e compila o binário
RUN go mod download
RUN go build -o /app/ordersystem main.go wire.go

# Etapa final
FROM debian:bullseye-slim

# Copia o binário e o arquivo .env do estágio de build
COPY --from=builder /app/ordersystem /usr/local/bin/ordersystem
COPY --from=builder /app/cmd/ordersystem/.env /usr/local/bin/.env

# Define o diretório de trabalho
WORKDIR /usr/local/bin

# Expõe as portas necessárias
EXPOSE 8000
EXPOSE 8080
EXPOSE 50051

# Comando para executar a aplicação
CMD ["ordersystem"]