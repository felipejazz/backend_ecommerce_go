# Use a imagem oficial do golang como base
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copie os arquivos necessários para o contêiner
COPY . .

# Instale as dependências do Go
RUN go mod download

# Comando para iniciar o container e aguardar
CMD ["tail", "-f", "/dev/null"]
