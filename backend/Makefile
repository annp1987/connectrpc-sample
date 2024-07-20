GOPATH:=$(shell go env GOPATH)
include env
.PHONY: env

format:
	gofmt -l -s -w .

analyzer:
	go vet ./...

test:
	go test -v ./... -cover

proto-build:
	protoc --proto_path=./proto --go_out=./proto  --go_opt paths=source_relative --connect-go_out=./proto --connect-go_opt paths=source_relative ./proto/greet.proto

docker:
	docker build -t sms-app:latest .

run: format analyzer
	go run ./cmd/main.go

.PHONY: format analyzer test docker run