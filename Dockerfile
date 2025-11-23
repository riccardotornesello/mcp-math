# Build stage
FROM golang:1.25 AS builder

WORKDIR /app

# Install modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o mcp-math .

# Final stage - use scratch for minimal image
FROM scratch

# Copy the binary from builder
COPY --from=builder /app/mcp-math /mcp-math

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/mcp-math"]
