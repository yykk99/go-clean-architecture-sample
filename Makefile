DOCKER_COMPOSE := docker-compose

.PHONY: up down logs app db

up:
	$(DOCKER_COMPOSE) up -d --build

down:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs -f --tail=100

app:
	$(DOCKER_COMPOSE) exec app sh

db:
	$(DOCKER_COMPOSE) exec db mysql -uroot -proot develop
