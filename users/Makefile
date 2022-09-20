ALL_TARGET = $(shell sed -n -e '/^$$/ { n ; /^[^ .\#][^ ]*:/ { s/:.*$$// ; p ; } ; }' $(MAKEFILE_LIST))
.PHONY: $$ALL_TARGET 
 
gen:
	# protoc pkg/pb/*.proto --go_out=plugins=grpc:.	
	protoc \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/protobuf/google/protobuf \
		-I pkg pkg/pb/*.proto --gofast_out=plugins=grpc:. 

test: down
	docker-compose up -d --force-recreate testdbuser	
	echo Starting for db...
	sleep 15
	docker-compose up migratedbuser && make quicktest && make down

quicktest: 	
	docker cp ./migrations/test/remove_user.sql testdbuser:/docker-entrypoint-initdb.d/dump.sql
	winpty docker exec -i -t testdbuser psql -d library -U db -f docker-entrypoint-initdb.d/dump.sql

	go test -v ./... -coverprofile cover.out

down:
	docker-compose down

buildapi:
	# docker-compose build --no-cache api
	docker-compose build api

runapi: 
	docker-compose up -d api

localapi:
	go run cmd/main.go

	