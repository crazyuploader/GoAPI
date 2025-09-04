#
# Created by Jugal Kishore -- 2025
#
FROM golang:1.25.1 AS builder

# Set work directory
WORKDIR /app

# Copy files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 \
    go build -a -installsuffix cgo \
    -ldflags='-w -s -extldflags "-static"' \
    -o api_server cmd/server/main.go

# Production stage
FROM alpine:3.22.1 AS runner

LABEL maintainer="Jugal Kishore <me@devjugal.com>"
LABEL org.opencontainers.image.source="https://github.com/crazyuploader/GoAPI"
LABEL org.opencontainers.image.description="Minimal Go API server using Fiber"

# Work directory
WORKDIR /app

# Copy the binary
COPY --from=builder /app/api_server .

# Create non-root user for building
RUN adduser -D -s /bin/sh -u 10001 appuser

# Set ownership and permissions
RUN chown appuser:appuser api_server \
    && chmod +x api_server

# Use non-root user
USER appuser

# Expose port
EXPOSE 3100

CMD ["./api_server"]
