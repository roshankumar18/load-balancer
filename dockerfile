    FROM golang:1.25-alpine AS builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go
    RUN CGO_ENABLED=0 GOOS=linux go build -o servers servers/servers.go


    # - Servers
    

    FROM alpine:3.18 AS  runner-loadbalancer
    WORKDIR /app
    COPY --from=builder /app/main .
    COPY configs ./configs

    EXPOSE 8080

    CMD ["./main"]