ALL_TARGET = $(shell sed -n -e '/^$$/ { n ; /^[^ .\#][^ ]*:/ { s/:.*$$// ; p ; } ; }' $(MAKEFILE_LIST))
.PHONY: $$ALL_TARGET 

pull_api: 
	docker-compose down
	docker pull maslow123/library-users
	docker pull maslow123/library-api-gateway

infratest: pull_api
	docker-compose up -d --force-recreate testdb
	echo Starting for db...
	sleep 15
	docker-compose up migratedb

test:
	cd users && make test
	cd api-gateway && make test

runapi: infratest
	docker-compose up -d --force-recreate userapi
	docker-compose up -d --force-recreate apigateway

down:
	docker-compose down

doc:
	cd api-gateway && make serve_swagger
