# Use the official Golang image
FROM golang:latest As builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN go mod download



FROM builder As final

# Build the Go application
RUN go build -o IMSController .


EXPOSE 10000

CMD ["./IMSController"]




