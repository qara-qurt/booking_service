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
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

# Stop test enviroment(postgres)
test_stop:
	docker-compose -f docker-compose.test.yml down

# Start application
run:
	docker-compose -f docker-compose.yml up --build

# Stop application
stop:
	docker-compose -f docker-compose.yml down 