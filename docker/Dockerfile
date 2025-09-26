FROM node:24.8.0-alpine AS node_builder

WORKDIR /build

COPY package.json package-lock.json ./
RUN npm install
COPY ./internal/interface/web ./internal/interface/web
RUN npm run tailwind

FROM golang:1.25.1 AS go_builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

FROM alpine:3.22.1

WORKDIR /app

COPY --from=go_builder /build/server .
COPY --from=node_builder /build/internal/interface/web/public ./internal/interface/web/public
EXPOSE 3000
CMD ["./server"]
