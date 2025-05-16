PROTO_DIR := proto
PROTO_SRC := $(shell find $(PROTO_DIR) -name '*.proto')
GO_OUT := shared/proto

generate-proto:
	rm -f $(GO_OUT)/*.go
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_OUT) --go-grpc_opt=paths=source_relative \
		$(PROTO_SRC)

run-all:
	docker compose -f compose.base.yaml up -d

stop-all:
	docker compose -f compose.base.yaml down

.PHONY: generate-proto run-all stop-all