
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/micro/proto/api/api.proto=github.com/micro/micro/v3/proto/api

.PHONY: proto
proto:
    
	protoc --openapi_out=. --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/posts.proto

.PHONY: docs
docs:
    
	protoc --openapi_out=. --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/posts.proto
	@redoc-cli bundle api-posts.json

.PHONY: build
build: proto

	go build -o post-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t post-service:latest
