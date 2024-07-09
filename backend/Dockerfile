FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod go.sum .env ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/
# COPY --from=builder /app/.env /app/
ENV APP_ENV=prod

CMD ["./main"]