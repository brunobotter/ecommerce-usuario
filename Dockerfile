FROM golang:1.21.4 as builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -o ecommerce-usuario

FROM scratch
COPY --from=builder /app/ecommerce-usuario /ecommerce-usuario
EXPOSE 8080
ENTRYPOINT ["/ecommerce-usuario"]

