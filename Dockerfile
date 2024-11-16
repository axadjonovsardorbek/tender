# File: Dockerfile
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application
RUN go build -o main cmd/main.go

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
