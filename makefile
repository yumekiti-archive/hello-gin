dc := docker-compose -f ./docker/docker-compose.yml

.PHONY: up
up:
	$(dc) up -d --build

.PHONY: down
down:
	$(dc) down

.PHONY: restart
restart:
	$(dc) restart

.PHONY: rm
rm:
	$(dc) down --rmi all

.PHONY: logs
logs:
	$(dc) logs -f

.PHONY: gin
gin:
	$(dc) exec /bin/sh gin

.PHONY: docker-rm
docker-rm:
	docker stop `docker ps -aq` ;
	docker rm `docker ps -aq`