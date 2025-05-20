# host ?= 35.240.184.74
# port ?= 54321
# user ?= postgres
# pass ?= 1
# db ?= konis_caffee

include .env

migrate\:init:
	PGPASSWORD=$(DB_PASSWORD) psql -U$(DB_USER) -d postgres -h $(DB_HOST) -p $(DB_PORT) -c "create database $(DB_NAME);"

migrate\:drop:
	PGPASSWORD=$(DB_PASSWORD) psql -U$(DB_USER) -d postgres -h $(DB_HOST) -p $(DB_PORT) -c "drop database if exists $(DB_NAME) with (force);"

migrate\:up:
	migrate -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path migrations up $(version)

migrate\:down:
	migrate -database postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable -path migrations down $(version)

migrate\:reset: 
	$(MAKE) migrate:drop user=$(DB_USER) db=$(DB_NAME)
	$(MAKE) migrate:init user=$(DB_USER) db=$(DB_NAME)
	$(MAKE) migrate:up user=$(DB_USER) pass=$(DB_PASSWORD) db=$(DB_NAME)


migrate\:create:
	migrate create -ext sql -dir migrations -seq create_$(table)_table

migrate\:alter:
	migrate create -ext sql -dir migrations -seq alter_$(table)_table

migrate\:insert:
	migrate create -ext sql -dir migrations -seq insert_$(table)_table


run:
	nodemon --exec go run main.go --signal SIGTERM