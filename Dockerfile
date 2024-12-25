FROM golang:latest AS builder

WORKDIR /app

COPY . .

EXPOSE 10000

RUN go mod download

# Build the application as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o IMSController .

FROM alpine:latest AS final

WORKDIR /app

COPY --from=builder /app/IMSController .

# Ensure the binary has execution permissions
RUN chmod +x IMSController


# Run the application
CMD ["./IMSController"]
