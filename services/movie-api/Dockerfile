FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.work .
COPY go.work.sum .
COPY pkg/go.mod pkg/go.mod

COPY services/movie-api/go.mod services/movie-api/go.mod
COPY services/auth-api/go.mod services/auth-api/go.mod
COPY services/reservation-api/go.mod services/reservation-api/go.mod

COPY services/movie-api/go.sum services/movie-api/go.sum

WORKDIR /app/services/movie-api
RUN go mod download

WORKDIR /app
COPY pkg/ pkg/
COPY services/movie-api/ services/movie-api/

# flag -o -> output path
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-bin services/movie-api/cmd/main.go

FROM alpine:latest

WORKDIR /root/

# In case we need external SSL conections
# RUN apk --no-cache add ca-certificates 

COPY --from=builder /api-bin .

EXPOSE 8080

CMD ["./api-bin"]