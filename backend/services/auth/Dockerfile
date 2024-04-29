FROM golang:1.22.2

WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o main ./cmd/app/main.go

# Command to run the service
EXPOSE 8082
CMD ["./main"]
