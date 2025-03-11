# main - Makefile for servers-mcppc-config
.PHONY: help build test clean deploy

# Variables
APP_NAME=servers-mcppc-config
VERSION=1.0.0

help:
	@echo "Available commands:"
	@echo "  build    - Build the application"
	@echo "  test     - Run tests"
	@echo "  clean    - Clean build artifacts"
	@echo "  deploy   - Deploy the application"

build:
	@echo "Building $(APP_NAME) v$(VERSION)"
	npm run build
	@echo "Build completed successfully"

test:
	@echo "Running tests..."
	npm test
	@echo "Tests completed"

clean:
	@echo "Cleaning build artifacts..."
	rm -rf dist/ node_modules/ coverage/
	@echo "Clean completed"

deploy:
	@echo "Deploying $(APP_NAME)..."
	npm run deploy
	@echo "Deployment completed"

# Development
dev:
	npm run dev

# Production
start:
	npm start

# Code Update 1760952379-24314

# Code Update 1760952379-31539

# Code Update 1760952379-3301

# Additional Implementation 1760952379

# Touch update: 1760952381
