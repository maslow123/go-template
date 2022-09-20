ALL_TARGET = $(shell sed -n -e '/^$$/ { n ; /^[^ .\#][^ ]*:/ { s/:.*$$// ; p ; } ; }' $(MAKEFILE_LIST))
.PHONY: $$ALL_TARGET 
 
gen:
	protoc \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2 \
		-I D:/me/exercise/go/pkg/mod/github.com/gogo/protobuf@v1.3.2/gogoproto \
		-I pkg pkg/users/pb/*.proto --gofast_out=plugins=grpc:. 

test:	
	docker cp ./migrations/test/remove_user.sql testdb:/docker-entrypoint-initdb.d/dump.sql
	winpty docker exec -i -t testdb psql -d library -U db -f docker-entrypoint-initdb.d/dump.sql

	go test -v ./... -coverprofile cover.out

down:
	docker-compose down

buildapi:
	docker-compose build --no-cache api-gateway

runapi: 
	docker-compose up -d api-gateway

runlocalapi:
	go run cmd/main.go

check_swagger:
	which swagger || GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve_swagger: gen swagger
	swagger serve -F=swagger swagger.yaml

