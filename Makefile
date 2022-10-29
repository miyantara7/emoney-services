build: setup
	@echo "--- Building binary file ---"
	@go build -o ./main server/grpc/main.go

grpc:
	@echo "--- running gRPC server in dev mode ---"
	@go run server/grpc/main.go

tidy:
	@go mod tidy

setup:
	@echo " --- Setup and generate configuration --- "
	@cp config/example/mysql.yml.example config/db/mysql.yml
	@cp config/example/server.yml.example config/server/server.yml
	@cp config/example/client.yml.example config/client/client.yml

build-docker: build
	@docker build --tag emoney-services .

protoc-docker:
	@docker container create --name emoney-services -p 9903:9903/tcp emoney-services