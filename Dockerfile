# Use the official Golang image as the base image
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o myapi ./cmd

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapi .

# Expose the port your application listens on
EXPOSE 8081

# Command to run the application
CMD ["./myapi"]