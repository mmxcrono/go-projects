.PHONY: all run

# Include the .env file
include .env
export

build:
	go build -o myapp cmd/main.go

up:
	docker run \
	-e DB_USER=${DB_USER} \
	-e DB_PASS=${DB_PASS} \
	-e DB_ADDRESS=${DB_ADDRESS} \
	-e DB_NAME=${DB_NAME} \
	--rm \
	--name sql-api \
	sql-api