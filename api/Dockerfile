# Use a imagem oficial do Golang como imagem base
FROM golang:1.21.5

# Define o diretório de trabalho dentro do container
WORKDIR /go/src/app

# Copia os arquivos da aplicação Go para o container
COPY . .

# Baixa as dependências da aplicação Go (se houver)
RUN go mod download

# Compila a aplicação Go
RUN go build -o /go/bin/app

# Define o comando para rodar a aplicação
CMD ["/go/bin/app"]
