FROM golang:1.22-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/backend /app/backend

EXPOSE 8888

ENTRYPOINT ["/app/backend"]


# Normal
# FROM golang:1.23.0

# WORKDIR /app 

# COPY . /app/

# RUN go mod tidy

# RUN go build -o backend

# EXPOSE 8080

# ENTRYPOINT ["/app/backend"]