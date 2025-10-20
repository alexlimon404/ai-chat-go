.PHONY: run build clean test deps migrate

# Run the application
run:
	@if [ ! -f .env ]; then cp .env.example .env; fi
	go run main.go

# Build the application
build:
	go build -o ai-chat-go main.go

# Clean build artifacts
clean:
	rm -f ai-chat-go

# Run tests
test:
	go test -v ./...

# Install dependencies
deps:
	go mod download
	go mod tidy

# Database migration (auto-runs on app start)
migrate:
	go run main.go
