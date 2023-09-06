APP             = g-api
SERVER_BIN  	= ./cmd/$(APP)/$(APP)
FOLDER          = ./app ./cmd ./core ./pkg
ALL_PATH        = ./...

UNPACK_PATH = $$path

DOCKER_CMD        = docker
DOCKER_BUILD      = $(DOCKER_CMD) build
DOCKER_PUSH       = $(DOCKER_CMD) push
DOCKER_IMAGE_NAME = g-api



i:
	@go mod download
test:
	@go test -v $(ALL_PATH) -cover
build:
	@go build -o $(SERVER_BIN) ./cmd/$(APP)
clean:
	@go clean
	rm -rf $(SERVER_BIN)
pack:
	tar -cvzf $(APP).tar.gz $(APP) $(FOLDER)
unpack:
	tar -zxf $(APP).tar.gz -C $(UNPACK_PATH)
docker_build:
	@echo "開始打包 Docker Image - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_BUILD) -t $(DOCKER_IMAGE_NAME) .
docker_push:
	@echo "開始 push docker image - $(DOCKER_FULL_IMAGE)"
	$(DOCKER_PUSH) $(DOCKER_IMAGE_NAME)
swag:
	@echo "初始化swag"
	@swag init --generalInfo ./cmd/$(APP)/main.go --output ./core/swagger