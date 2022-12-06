APP_NAME=bookstorecrud

all_docker: lint-client lint-server up up-elastic
all: lint-server lint-server up-elastic runS runC

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
	docker-compose -f docker-compose.yml up

start: ## Start running containers
	docker-compose start

stop: ## Stop running containers
	docker-compose stop

rm: stop ## Stop and remove running containers
	docker rm $(APP_NAME)

lint-server: ## Run golangci-lint on Server
	cd server; \
	golangci-lint run
	cd server; \
	go vet ./...
	@echo "Golangci-lint and vet tests are finished successful"

lint-client: ## Run golangci-lint on Client
	cd client; \
	golangci-lint run
	cd client; \
	go vet ./...
	@echo "Golangci-lint and vet tests are finished successful"

create-elastic:
	docker network create elastic

run-elastic: ##Run elasticsearch container
	docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2

run-kibana: ## Run Kibana container
	docker run --name kib-01 --net elastic -p 5601:5601 docker.elastic.co/kibana/kibana:7.4.0

up-elastic:
	docker-compose -f docker-compose.elastic.yml up

get-grpc:
	go get google.golang.org/protobuf/cmd/protoc-gen-go \
             google.golang.org/grpc/cmd/protoc-gen-go-grpc

mock_gen:
	go generate ./...


MOCKS_DESTINATION=mocks
.PHONY: mocks
# put the files with interfaces you'd like to mock in prerequisites
# wildcards are allowed
mocks: server/service/service.go server/repository/repository.go
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done

