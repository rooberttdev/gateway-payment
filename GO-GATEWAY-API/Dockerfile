# Dockerfile
FROM golang:alpine

WORKDIR /app

# Copia arquivos de dependência primeiro (para cache)
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copia o restante da aplicação
COPY . .

# Comando padrão do container
CMD ["go", "run", "cmd/app/main.go"]
