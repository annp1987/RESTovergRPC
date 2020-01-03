GRPC_GATEWAY_DIR := $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway 2> /dev/null)
GO_INSTALLED := $(shell which go)
PROTOC_INSTALLED := $(shell which protoc)
BINDATA_INSTALLED := $(shell which go-bindata 2> /dev/null)
PGGG_INSTALLED := $(shell which protoc-gen-grpc-gateway 2> /dev/null)
PGG_INSTALLED := $(shell which protoc-gen-go 2> /dev/null)
SERVER_IMAGE_NAME     ?= server:latest
SERVER_DOCKER_PATH    ?= ./docker/server/Dockerfile
DB_IMAGE_NAME     ?= db:latest
DB_DOCKER_PATH    ?= ./docker/postgres/Dockerfile


install-tools:
ifndef GO_INSTALLED
	$(error "go is not installed)
endif
ifndef PROTOC_INSTALLED
	$(error "protoc is not installed)
endif
ifndef BINDATA_INSTALLED
	go get -u github.com/kevinburke/go-bindata/go-bindata@master
endif
ifndef PGGG_INSTALLED
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
endif
ifndef PGG_INSTALLED
	go get -u github.com/golang/protobuf/protoc-gen-go
endif

gen-proto:
	protoc -I. \
    	   -I$(GRPC_GATEWAY_DIR)/third_party/googleapis \
    	   --go_out=plugins=grpc:directory \
    	   --grpc-gateway_out=logtostderr=true:directory \
    	   --proto_path directory directory.proto

build:
	go build -o ./docker/server/server main.go

images:
	docker build -t $(SERVER_IMAGE_NAME) $(SERVER_DOCKER_PATH)
	docker build -t $(DB_IMAGE_NAME) $(DB_DOCKER_PATH)
