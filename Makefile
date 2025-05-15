PROTO_DIR := proto
PROTO_SRC := $(shell find $(PROTO_DIR) -name '*.proto')
GO_OUT := .

generate-proto:
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GO_OUT) \
		--go-grpc_out=$(GO_OUT) \
		$(PROTO_SRC)

run-all:
	docker compose -f compose.base.yaml up -d

stop-all:
	docker compose -f compose.base.yaml down

.PHONY: generate-proto run-all stop-all