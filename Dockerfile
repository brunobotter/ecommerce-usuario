# Etapa de build
FROM golang:1.22.3 AS builder

# Configura o diretório de trabalho dentro do container
WORKDIR /app

# Copia o go.mod e go.sum e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte
COPY . .

# Compila o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ecommerce-usuario .

# Etapa de produção
FROM scratch

# Copia o binário compilado da etapa de build
COPY --from=builder /app/ecommerce-usuario /ecommerce-usuario

# Define a porta que o aplicativo vai expor
EXPOSE 8080

# Define o ponto de entrada do container
ENTRYPOINT ["/ecommerce-usuario"]