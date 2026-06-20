FROM golang:1.26.4-alpine AS builder
WORKDIR /services/auth-service
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM gcr.io/distroless/static-debian12
WORKDIR /services/auth-service
COPY --from=builder /services/auth-service/app app
CMD [ "./app" ]