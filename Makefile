include envs/dev/.env

.PHONY: run
.DEFAULT_GOAL := run


build: swagger
	go build -v ./cmd/app

run: swagger
	go run ./cmd/app/main.go

swagger: tidy
	swag init -g cmd/app/main.go
	swag fmt

test: tidy
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/go_auth_users_test?sslmode=disable" up
	go test -v -cover -race -timeout 30s ./...
	
testcover: tidy
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/go_auth_users_test?sslmode=disable" up
	go test -cover -coverpkg ./... ./...

tidy:
	go mod tidy

lint:
	golangci-lint run ./...
	
# --fast  --config=./.golangci.yaml

################
## MIGRATIONS ##
################

makemigrations:
	migrate create -ext sql -dir migrations $(name)

migratetables:
	migrate -path migrations -database "postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" $(mode)

################
#### DOCKER ####
################

runstagebuild:
	docker-compose -f docker-compose.stage.yaml -p="go-auth-stage" up --build

runstage:
	docker-compose -f docker-compose.stage.yaml -p="go-auth-stage" up

up:	swagger
	docker-compose -p="go-auth-prod" up -d --build 
