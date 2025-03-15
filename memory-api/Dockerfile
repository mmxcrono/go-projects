# ---- Stage 1: Build ----
FROM golang:latest AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application (static binary recommended)
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# ---- Stage 2: Run ----
FROM alpine:latest

# Set a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy the built binary from the previous stage
COPY --from=builder /app/app .

# Grant executable permissions
RUN chmod +x app

# Set user permissions
USER appuser

# Expose application port (if applicable)
EXPOSE 8080

# Run the application
CMD ["./app"]
