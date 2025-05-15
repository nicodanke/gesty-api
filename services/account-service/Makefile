include .env

migrate-up:
	migrate -path db/migrations -database ${DSN} -verbose up

migrate-up1:
	migrate -path db/migrations -database ${DSN} -verbose up 1

migrate-down:
	migrate -path db/migrations -database ${DSN} -verbose down

migrate-down1:
	migrate -path db/migrations -database ${DSN} -verbose down 1

sqlc-generate:
	sqlc generate

server:
	go run main.go

.PHONY: migrate-up migrate-up1 migrate-down migrate-down1 server sqlc-generate