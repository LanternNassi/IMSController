# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Copy source code
COPY . .

# Build-time environment variables
ARG DBHOST
ARG DBPORT
ARG DBUSER
ARG DBPASSWORD
ARG DBNAME
ARG APPHOST=0.0.0.0:10000

ARG test_DBHOST=db
ARG test_DBPORT=5432
ARG test_DBUSER=postgres
ARG test_DBPASSWORD=postgres
ARG test_DBNAME=testdb

# Create .env file with the required values
RUN echo "DBHOST=${DBHOST}" >> .env && \
    echo "DBPORT=${DBPORT}" >> .env && \
    echo "DBUSER=${DBUSER}" >> .env && \
    echo "DBPASSWORD=${DBPASSWORD}" >> .env && \
    echo "DBNAME=${DBNAME}" >> .env && \
    echo "APPHOST=${APPHOST}" >> .env && \
    echo "test_DBHOST=${test_DBHOST}" >> .env && \
    echo "test_DBPORT=${test_DBPORT}" >> .env && \
    echo "test_DBUSER=${test_DBUSER}" >> .env && \
    echo "test_DBPASSWORD=${test_DBPASSWORD}" >> .env && \
    echo "test_DBNAME=${test_DBNAME}" >> .env

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
