proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	@ if ! which protoc-gen-go > /dev/null; then \
		echo "error: protoc-gen-go not installed" >&2; \
		exit 1; \
	fi

	@ echo Compiling Protobufs
	protoc \
		--go_out=Mgrpc/service_config/service_config.proto=/proto-gen/api:. \
		--go-grpc_out=Mgrpc/service_config/service_config.proto=/proto-gen/api:. \
		diago-idl/proto/worker.proto


.PHONY: build
build:
	GOOS=linux go build cmd/main.go
	docker build -f build/Dockerfile -t diago-worker .

run:
	@ go run main

local:
	go build cmd/main.go
