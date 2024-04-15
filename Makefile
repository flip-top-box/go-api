build:
	@go build -o bin/GOAPI cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/GOAPI

.PHONY: migrate-up migrate-down

DB_MIGRATE_PATH=./cmd/migrate/main.go  # Adjust this path to where your main.go for migration is located
MIGRATION_EXECUTABLE=./cmd/migrate/migrations/migrate

build-migrate:
	@echo "Building migration tool..."
	@go build -o $(MIGRATION_EXECUTABLE) $(DB_MIGRATE_PATH)

migrate-up: build-migrate
	@echo "Applying database migrations..."
	@$(MIGRATION_EXECUTABLE) up
	@echo "Migrations applied successfully."

migrate-down: build-migrate
	@echo "Reverting database migrations..."
	@$(MIGRATION_EXECUTABLE) down
	@echo "Migrations reverted successfully."

clean:
	@echo "Cleaning up..."
	@rm -f $(MIGRATION_EXECUTABLE)
	@rm -f bin/GOAPI
