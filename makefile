include .$(PWD)/.env

create-app:
	docker-compose up -d

restart-app:
	docker-compose up -d --build