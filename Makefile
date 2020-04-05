.DEFAULT_GOAL := docker-run
APP_NAME=hellogrpc

.PHONY: docker-build
docker-build:
	docker build --build-arg APP_NAME=$(APP_NAME) -t $(APP_NAME) .

.PHONY: docker-run
docker-run:
	docker run -p 9090:9090 $(APP_NAME)

.PHONY: hellogrpc
hellogrpc:
	go install github.com/thepatrik/hellogrpc
