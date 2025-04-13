APP_NAME = go-backend
SERVICE_NAME = backend
DOCKER_COMPOSE = docker-compose

.PHONY: all build up down restart logs clean

all: build

build:
	docker build -t $(APP_NAME) .

up:
	$(DOCKER_COMPOSE) up -d --build

down:
	$(DOCKER_COMPOSE) down

restart:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) up -d --build

logs:
	$(DOCKER_COMPOSE) logs -f $(SERVICE_NAME)

clean:
	$(DOCKER_COMPOSE) down -v --rmi all --remove-orphans
