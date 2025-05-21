# Start from official Go image (use slim for smaller size)
FROM golang:1.24.3

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy rest of the code
COPY . .

# Build the binary inside the container
RUN go build -o main .

# Expose port your app listens on
EXPOSE 3001

# Command to run the binary
CMD ["./main"]
