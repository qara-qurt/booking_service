# Migrate
create-migration:
	migrate create -ext sql -dir ./internal/repository/postgres/migrations -seq init

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations down

migrate-drop:
	migrate -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -path ./internal/repository/postgres/migrations drop


# Start test enviroment(postgres and app)
# Build docker-compose.test.yml to create a test environment with a PostgreSQL database and app
# Run the migrations for app docker container
# Run the tests in the app docker container
# Tear down the test environment after the tests are done
test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit


# Stop test enviroment(postgres)
test_stop:
	docker-compose -f docker-compose.test.yml down

# Start application with docker-compose
run:
	docker-compose -f docker-compose.yml up --build

# Stop application
stop:
	docker-compose -f docker-compose.yml down 