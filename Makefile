DB_DSN := postgres://postgres:7ab816zAd@localhost:15432/postgres?sslmode=disable

MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up

migrate-down: 
	$(MIGRATE) down

run: 
	go run cmd/main.go