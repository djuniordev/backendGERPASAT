# Etapa 1: Usar a imagem oficial do Go para compilar a aplicação
FROM golang:1.20-alpine AS build

# Instalar dependências necessárias para o Air
RUN apk update && apk add --no-cache bash git

# Criar um diretório para a aplicação
WORKDIR /app

# Copiar o arquivo go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar as dependências do Go
RUN go mod tidy

# Copiar o restante do código da aplicação para o contêiner
COPY . .

# Instalar o Air
RUN go install github.com/cosmtrek/air@latest

# Etapa 2: Configuração do ambiente para execução da aplicação com Air
FROM golang:1.20-alpine

# Instalar dependências do sistema (caso precise de algo mais, como o bash ou outros pacotes)
RUN apk update && apk add --no-cache bash

# Criar um diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos da etapa de build (binário e código) para a nova imagem
COPY --from=build /app /app

# Expor a porta em que o servidor Go vai rodar
EXPOSE 8080

# Definir o comando padrão para rodar o Air
CMD ["air"]
