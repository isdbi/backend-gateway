proto-generate:
	protoc --go_out=. --go-grpc_out=. --twirp_out=. pkg/pb/service.proto

run:
	go run cmd/api/main.go

generate:
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc generate

all: proto-generate generate

