# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git protobuf-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Generate protobuf files
RUN make generate-proto

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api ./services/api

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/api .

# Copy any necessary config files
COPY --from=builder /app/compose.base.yaml .
COPY --from=builder /app/db_password.txt .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./api"] 