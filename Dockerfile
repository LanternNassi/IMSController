# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Copy source code
COPY . .

# Download dependencies
RUN go mod download

# Build the application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o IMSController .

# Stage 2: Final image
FROM alpine:latest AS final

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/IMSController .

# Ensure the binary has execution permissions
RUN chmod +x IMSController

# Expose the application port
EXPOSE 10000

# Run the application
CMD ["./IMSController"]
