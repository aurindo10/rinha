# Use uma imagem oficial do Go para compilar o aplicativo
FROM golang:1.18 AS build
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app .

# Use uma imagem mais leve para rodar o aplicativo
FROM alpine:latest
COPY --from=build /app /app
CMD ["/app"]
