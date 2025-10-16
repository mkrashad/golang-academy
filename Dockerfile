# Base the image on the official Go image
FROM golang:1.21.0-alpine AS builder

# Create the application directory
WORKDIR /app

# Enable Go modules
ENV GO111MODULE=on

# Copy and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the application source
COPY . .

# Build the application
RUN go build -o main /app/cmd/server/main.go

# Execution stage
FROM scratch

WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/main .

# Expose the port
EXPOSE 8081

# Execute the application command
CMD ["./main"]