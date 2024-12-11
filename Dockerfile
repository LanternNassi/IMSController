# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app


# Copy source code
COPY . .


# Download dependencies
RUN go mod download

# Build the application
RUN go build -o IMSController .

# Stage 2: Final image
FROM alpine:latest AS final

WORKDIR /app

# Copy the compiled binary and .env file from the builder stage
COPY --from=builder /app/IMSController .
COPY --from=builder /app/.env .

# Expose the application port
EXPOSE 10000

# Run the application
CMD ["./IMSController"]
