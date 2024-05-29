
# Application
project := "story-tracker"
version := "0.1.0"

# Build
output_dir :="bin"
binary_name := project
binary_path :="." / output_dir / binary_name


# Build the application
build:
	@echo "Building "{{ binary_name }}"..."
	@sqlc generate
	@go build -v -o {{ binary_path }} cmd/api/main.go

# Run the application
run:
	@echo "Running "{{ binary_name }}"..."
	@go run cmd/api/main.go

# Run test suite
test:
	@echo "Running tests..."
	@go test -v ./...

# Run linters
lint:
	@go vet ./...
	@sqlc vet

# Run formatting
fmt:
	@go fmt ./...

# Clean the binary
clean:
	@echo "Cleaning project..."
	@rm -rf ./{{ output_dir }}

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Live reload enabled...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# Generate code from SQL
sqlc:
	@echo "Generating code from SQL..."
	@sqlc generate

# Create DB container
db-up:
	@echo "Starting database instance..."
	@docker compose --env-file .env.local up

# Shutdown DB container
db-down:
	@echo "Shutting down database..."
	@docker compose down;

# Execute database migrations
migrate-up:
	@echo "Running database migrations..."
	@dbmate --env-file ".env.local" --migrations-dir "db/schema" --migrations-table "migrations" up

# Revert the last database migration
migrate-down:
	@echo "Reverting last database migration..."
	@dbmate --env-file ".env.local" --migrations-dir "db/schema" --migrations-table "migrations" down

