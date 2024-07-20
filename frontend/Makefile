docker:
	docker build -t connect-fe:latest .

proto-build:
	protoc --proto_path=./src/proto --es_out=./src/proto --connect-es_out=./src/proto ./src/proto/greet.proto

.PHONY: docker proto-build