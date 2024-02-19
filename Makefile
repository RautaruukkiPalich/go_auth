.PHONY: run
.DEFAULT_GOAL := run

# run: swagger
# 	go run ./cmd/app/main.go

# build: swagger
# 	go build -v ./cmd/app

# swagger: tidy
# 	swag init
# 	swag fmt


run: tidy
	go run ./cmd/app/main.go

test: tidy
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/go_auth_users_test?sslmode=disable" up
	go test -v -race -timeout 30s ./...

tidy:
	go mod tidy

makemigrations:
	migrate create -ext sql -dir migrations $(name)

migratetables:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/go_auth_users?sslmode=disable" $(mode)