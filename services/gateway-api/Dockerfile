FROM golang:1.26.4-alpine AS builder
WORKDIR /services/gateway-api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM gcr.io/distroless/static-debian12
WORKDIR /services/gateway-api
COPY --from=builder /services/gateway-api/app app
CMD [ "./app" ]