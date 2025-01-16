include .env

BIN_DIR:=$(CURDIR)/../bin

generate-user-api:
	mkdir -p pkg/user_api_v1
    protoc --proto_path api/user_api_v1 \
   	--go_out=pkg/user_api_v1 --go_opt=paths=source_relative \
    --plugin=protoc-gen-go=../../bin/protoc-gen-go \
    --go-grpc_out=pkg/user_api_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=../../bin/protoc-gen-go-grpc \
	api/user_api_v1/user.proto

migration-status:
	$(BIN_DIR)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-up:
	$(BIN_DIR)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	$(BIN_DIR)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v