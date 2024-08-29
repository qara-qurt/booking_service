# Migrate
create-migration:
	migrate create -ext sql -dir ./internal/repository/postgres/migrations -seq init

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations down

migrate-drop:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations drop


# Start test enviroment(postgres)
test:
	docker-compose -f docker-compose.test.yml up --build -d
	@echo "Waiting for the database to be ready..."
	@sleep 5  # Wait for 5 seconds to give PostgreSQL time to initialize
	migrate -path ./internal/repository/postgres/migrations -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable up
	go test ./...
	docker-compose -f docker-compose.test.yml down


# Stop test enviroment(postgres)
test_stop:
	docker-compose -f docker-compose.test.yml down

# Start application
run:
	docker-compose -f docker-compose.yml up --build

# Stop application
stop:
	docker-compose -f docker-compose.yml down 