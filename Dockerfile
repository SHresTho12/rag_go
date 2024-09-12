# Stage 1: Build the Go application
FROM golang:1.23-alpine as build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app (output binary will be named "app")
RUN go build -o rag_go .

# Stage 2: Run the Go application
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the pre-built binary from the build stage
COPY --from=build /app/rag_go .

# Expose port (change if your app uses a different port)
EXPOSE 8080

# Command to run the application
CMD ["./rag_go"]
