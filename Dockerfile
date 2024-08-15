# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o IMSController .

# Expose the port the application runs on
EXPOSE 10000

# Command to run the application
CMD ["./IMSController"]