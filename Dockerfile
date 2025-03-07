# Use Go 1.24 bookworm as base image
FROM golang:1.24-alpine AS base

# Install necessary dependencies (SQLite client or runtime, etc.)
RUN apk add --no-cache sqlite

# Move to working directory /app
WORKDIR /app

# Copy the go.mod and go.sum files to the /app directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app. The binary will be output as 'main' (you can change this to whatever name you prefer)
RUN GOOS=linux go build -o /app/main /app/cmd/api/main.go

# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/main"]
