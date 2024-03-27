# Simple Makefile for a Go project

# Build the application
all: build

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@npm install
	@cp ./node_modules/htmx.org/dist/htmx.min.js ./cmd/web/static/htmx.min.js
	@cp ./node_modules/htmx.org/dist/ext/ws.js ./cmd/web/static/htmx-ws.js
	@cp ./node_modules/alpinejs/dist/cdn.min.js ./cmd/web/static/alpine.cdn.min.js

buildDeps:
	@echo "Building dependencies..."
	@npm run build:css
	@templ generate

build: buildDeps
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run: buildDeps
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	-rm -f main
	-rm -rf node_modules
	-rm ./cmd/web/static/built-styles.css
	-rm ./cmd/web/static/htmx-ws.js
	-rm ./cmd/web/static/htmx.min.js

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean install
