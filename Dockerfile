# --- Builder Stage: Build Go binary ---
FROM golang:1.23-bullseye AS builder
RUN apt-get update && apt-get upgrade -y

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
RUN go build -o discord-bot ./

# --- Final Stage: Playwright + Go binary ---
FROM mcr.microsoft.com/playwright:v1.51.1-jammy

WORKDIR /app

COPY --from=builder /app/discord-bot /app/discord-bot

# Environment vars with placeholders
ENV TEXT_CHAN_ID ""
ENV GUILD_ID ""
ENV TOKEN ""
ENV BOT_PREFIX ""
ENV LOG_LEVEL "INFO"

# Command to run the Discord bot
CMD ["./discord-bot"]
