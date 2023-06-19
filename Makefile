server:
	go run cmd/main.go

migrate:
	go run migrations/main.go

docs:
	swag init --parseDependency -g cmd/main.go