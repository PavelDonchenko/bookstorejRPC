APP_NAME=bookstorecrud

create_proto:
	cd server; \
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	cd server; \
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/

clean_proto:
	cd server; \
	rm gen/proto/*.go

runC:
	cd client; \
	go run main.go

runS:
	cd server; \
	go run main.go

build: ## Build the release and develoment container. The development
	docker-compose build

dev: ## Run container in development mode
	docker-compose build --no-cache $(APP_NAME) && docker-compose run $(APP_NAME)

# Build and run the container
up: ## Spin up the project
	docker-compose up

start: ## Start running containers
	docker-compose start

stop: ## Stop running containers
	docker-compose stop

rm: stop ## Stop and remove running containers
	docker rm $(APP_NAME)

lintS: ## Run golangci-lint on Server
	cd server; \
	golangci-lint run
	cd server; \
	go vet ./...
	echo "Golangci-lint and vet tests are finished successful"

lintC: ## Run golangci-lint on Client
	cd client; \
	golangci-lint run
	cd client; \
	go vet ./...
	echo "Golangci-lint and vet tests are finished successful"
