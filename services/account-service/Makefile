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

generate-grpc-gateway:
	rm -f gw/*.go
	rm -f doc/swagger/*.swagger.json
	protoc -I . \
		--proto_path=../../proto \
		--grpc-gateway_out ./gw \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=inventApp \
		--grpc-gateway_opt standalone=true \
		--grpc-gateway_opt grpc_api_configuration=account-service.yaml \
		../../proto/account-service/account_service.proto
	statik -src=./doc/swagger -dest=./doc

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

.PHONY: migrate-up migrate-up1 migrate-down migrate-down1 server sqlc-generate generate-grpc-gateway db_docs db_schema