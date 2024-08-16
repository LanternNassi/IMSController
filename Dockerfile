# Use the official Golang image
FROM golang:latest As builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN go mod download

# Build the Go application
RUN go build -o IMSController .


FROM builder As final

EXPOSE 10000

CMD ["./IMSController"]




