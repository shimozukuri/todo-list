FROM golang:1.27.0-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/todo-list/exe /app/cmd/todo-list

FROM alpine:3.24
WORKDIR /app
COPY --from=builder /app/cmd/todo-list/exe /app
CMD ["/app/exe"]