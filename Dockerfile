# Step 1: Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go modules files & download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /worker ./worker/worker.go

RUN chmod +x /main /worker

# Step 2: Runtime Stage
# FROM alpine:latest

# WORKDIR /root/

# Install dependencies (optional, for debugging)
# RUN apk --no-cache add ca-certificates

# Copy the compiled binary from builder stage
# COPY --from=builder /app/main .

# Expose the application's port
EXPOSE 8080:8080

# Run the application
CMD ["/main"]
