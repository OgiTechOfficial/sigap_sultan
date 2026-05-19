FROM golang:1.20-alpine

WORKDIR /app

# Install git and other build tools
RUN apk update && apk add --no-cache git

COPY go.mod go.sum ./
ENV GOPROXY=direct
RUN go mod download

COPY . .

# Expose port (adjust if necessary, usually Fiber runs on 8080 or what's in .env)
EXPOSE 8080

# Create logs directory
RUN mkdir -p logs

CMD ["go", "run", "main.go"]
