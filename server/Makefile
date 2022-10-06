APP_NAME=bookstorecrud

create_proto:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/

clean_proto:
	rm gen/proto/*.go
build: ## Build the release and develoment container. The development
	docker-compose build

dev: ## Run container in development mode
	docker-compose build --no-cache $(APP_NAME) && docker-compose run $(APP_NAME)

# Build and run the container
up: ## Spin up the project
	docker-compose up

stop: ## Stop running containers
	docker stop

rm: stop ## Stop and remove running containers
	docker rm $(APP_NAME)

lint: ## Run golangci-lint
	golangci-lint run
	go vet ./...
	echo "Golangci-lint and vet tests are finished successful"
