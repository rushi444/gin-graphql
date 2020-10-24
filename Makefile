gqlgen-init:
	go run github.com/99designs/gqlgen init

gql-generate:
	go run github.com/99designs/gqlgen generate

d-build:
	docker-compose build

d-up: 
	docker-compose up

dock: d-build d-up

run:
	go run main.go