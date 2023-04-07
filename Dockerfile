  # Step 1: Dependecies caching
FROM golang:1.20.1-alpine3.17 AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

  # Step 2: Builder
FROM golang:1.20.1-alpine3.17 AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/app

  # Step 3: Final
FROM scratch
COPY --from=builder /bin/app /app
CMD ["/app"]
