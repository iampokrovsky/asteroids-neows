db := "postgres://admin:admin@localhost:5432/orders_db?sslmode=disable"

migrate_up:
	migrate -path ./migrations -database $(db) up

migrate_down:
	migrate -path ./migrations -database $(db) down

run:
	go run ./cmd/app

run_pub:
	go run ./integration-test