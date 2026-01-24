FROM golang:1.25.5 AS builder

WORKDIR /app

# Copy workspace definitions
COPY go.work .
COPY go.work.sum .

# Copy go.mod from ALL services
COPY pkg/go.mod pkg/go.mod
COPY services/movie-api/go.mod services/movie-api/go.mod
COPY services/auth-api/go.mod services/auth-api/go.mod
COPY services/reservation-api/go.mod services/reservation-api/go.mod

# Copy current go.sum
COPY services/auth-api/go.sum services/auth-api/go.sum

WORKDIR /app/services/auth-api
RUN go mod download

WORKDIR /app
COPY pkg/ pkg/
COPY services/auth-api/ services/auth-api/

# Compile
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-bin services/auth-api/cmd/main.go

FROM alpine:latest

WORKDIR /root/

# In case we need external SSL conections
# RUN apk --no-cache add ca-certificates 

COPY --from=builder /api-bin .

EXPOSE 8080

CMD ["./api-bin"]