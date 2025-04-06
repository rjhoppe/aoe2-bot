# Use an official Go runtime as a parent image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download -x

# Copy the entire project source code
COPY . .

# Build the Go application
RUN go build -o discord-bot ./

# --- Final Stage: Create a smaller runtime image ---
FROM alpine:latest

# Install necessary runtime dependencies (if any)
# For a simple bot, you might not need extra dependencies
# RUN apk add --no-cache <your-runtime-dependencies>

# Set the working directory
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/discord-bot /app/discord-bot

# Environment vars with placeholders
ENV TEXT_CHAN_ID ""
ENV GUILD_ID ""
ENV TOKEN ""
ENV BOT_PREFIX ""
ENV LOG_LEVEL "INFO"

# Command to run the Discord bot
CMD ["./discord-bot"]