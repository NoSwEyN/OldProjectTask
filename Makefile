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

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
lint:
	golangci-lint run 

users:
	oapi-codegen -config openapi/.openuser -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/user.gen.go

