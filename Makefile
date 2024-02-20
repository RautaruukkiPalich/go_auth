.PHONY: run
.DEFAULT_GOAL := run

# build: swagger
# 	go build -v ./cmd/app


run: swagger
	go run ./cmd/app/main.go

swagger: tidy
	swag init -g cmd/app/main.go
	swag fmt

test: tidy
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/go_auth_users_test?sslmode=disable" up
	go test -v -race -timeout 30s ./...

tidy:
	go mod tidy


##############################################################################################################################

makemigrations:
	migrate create -ext sql -dir migrations $(name)

migratetables:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/go_auth_users?sslmode=disable" $(mode)