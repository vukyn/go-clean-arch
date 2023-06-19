server:
	go run cmd/main.go

migrate:
	go run migrations/main.go

gen-doc :
	swag init --parseDependency -g cmd/main.go