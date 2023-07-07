run:
	go run cmd/main.go

test:
	go test -cover ./...

migrate:
	go run migrations/main.go

tidy:
	go mod tidy

gen-doc:
	echo "Starting swagger generating"
	swag fmt
	swag init -g cmd/main.go

