# app commands
all: up

up:
	docker-compose build
	docker-compose up -d