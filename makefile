dc := docker-compose -f ./docker/docker-compose.yml

.PHONY: init
init:
	$(dc) up -d --build
	bash ./docker/mysql/sql.sh
	$(dc) exec -d gin /bin/sh -c "go run main.go"

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
	$(dc) exec gin /bin/sh

.PHONY: db
db:
	$(dc) exec db /bin/sh

.PHONY: docker-rm
docker-rm:
	docker stop `docker ps -aq` ;
	docker rm `docker ps -aq`