# ---- BUILD STAGE ----
FROM golang:1.24.1 AS builder

WORKDIR /app

# Copy go.mod and go.sum (if it exists)
COPY ../go.* ./

# Ensure go.mod exists before downloading dependencies
RUN test -f go.mod && go mod download || echo "go.mod not found!"

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY ../ .

# Build the Go application
RUN go build -o myapp .

# ---- RUN STAGE ----
FROM gcr.io/distroless/static:latest

WORKDIR /app

# Use a non-root user for security
USER 1001

# Copy the built binary from the builder stage
COPY --from=builder /app/myapp /app/

# Expose the application port (if needed)
EXPOSE 8080

# Start the application
CMD ["/app/myapp"]
