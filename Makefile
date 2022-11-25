# app commands
all: up

up:
	docker-compose build
	docker-compose up -d --scale server=5
down:
	docker-compose down

build:
	docker build -t drones-app .

remove:
	docker container rm drones-app

run: remove build
	docker run -d \
	-e POSTGRES_USER='linpostgres' \
	-e POSTGRES_DB='drones' \
	-e POSTGRES_PASSWORD='FrnEquyGW8t_30R4' \
	-e POSTGRES_HOST='lin-8444-1765-pgsql-primary.servers.linodedb.net' \
	-e POSTGRES_SSL='prefer' \
	--name drones-app \
	-p 6504:6504 drones-app

inline-cert:
	awk 'NF {sub(/\r/, ""); printf "%s\\n",$0;}' testing-ca-certificate.crt