# Etapa 1: Usar a imagem oficial do Go 1.22 para compilar a aplicação
FROM golang:1.22-alpine AS build

# Instalar dependências necessárias
RUN apk update && apk add --no-cache bash git

# Criar um diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o arquivo go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Rodar go mod tidy para garantir que as dependências estão corretas
RUN go mod tidy

# Baixar as dependências do Go
RUN go mod download

# Copiar o código fonte para o contêiner
COPY . .

# Compilar o código Go (apontando para o arquivo main.go dentro da pasta cmd)
RUN go build -o /app/meu-app ./cmd

# Etapa 2: Usar uma imagem mais leve para rodar a aplicação
FROM alpine:latest

# Instalar dependências mínimas para rodar o binário (como bash, se necessário)
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o binário gerado na etapa de build
COPY --from=build /app/meu-app /app/meu-app

# Expor a porta em que a aplicação Go vai rodar
EXPOSE 8080

# Comando para executar a aplicação
CMD ["/app/meu-app"]
