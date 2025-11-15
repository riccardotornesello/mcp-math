# Build stage
FROM golang:1.25 AS builder

WORKDIR /app

# Copy everything including vendor directory
COPY . .

# Build the application as a static binary using vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags="-w -s" -o mcp-math .

# Final stage - use scratch for minimal image
FROM scratch

# Copy CA certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder
COPY --from=builder /app/mcp-math /mcp-math

# Expose the port
EXPOSE 8080

# Run the application
CMD ["/mcp-math"]
