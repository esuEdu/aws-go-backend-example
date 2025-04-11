FROM golang:1.24.2

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o go-backend ./cmd/main.go

# Use a minimal base image for production
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=0 /app/go-backend .

# Command to run the executable
CMD ["./go-backend"]