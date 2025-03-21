.PHONY: all run

# Include the .env file
include .env
export

build-run: build run

build:
	docker build -t sql-api .

run:
	docker run \
	-e DB_USER=${DB_USER} \
	-e DB_PASS=${DB_PASS} \
	-e DB_ADDRESS=${DB_ADDRESS} \
	-e DB_NAME=${DB_NAME} \
	--rm \
	--name sql-api \
	sql-api

mysql:
	docker run \
	--name mysql \
	-e MYSQL_ROOT_PASSWORD=root \
	-e MYSQL_DATABASE=${DB_NAME} \
	-p 3306:3306 \
	-d \
	-v ./sql:/app/sql \
	-v ./mysql-data:/var/lib/mysql \
	--rm \
	mysql:latest
