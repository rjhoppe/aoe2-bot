# Use an official Go runtime as a parent image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install Playwright dependencies
RUN apk add --no-cache \
    nss \
    chromium \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    nodejs \
    npm \
    git

# Install Playwright CLI globally
RUN npm install -g playwright

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

# Set the working directory
WORKDIR /app

# Install Playwright dependencies in the runtime image too
RUN apk add --no-cache \
    nss \
    chromium \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    nodejs \
    npm

# Install Playwright CLI and browsers
RUN npm install -g playwright && playwright install --with-deps

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