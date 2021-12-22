SHELL := /bin/bash

# These will be provided to the target
GIT := github.com/gipsh/golang-rest-api-boilerplate/
VERSION := 1.0.0
BUILD := `git rev-parse HEAD`
AUTHOR := `whoami`
BUILD_DATE := `date +%Y%m%d%H%M%S`

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=$(GIT)build.Version=$(VERSION)\
				  -X=$(GIT)build.Time=$(BUILD_DATE)\
				  -X=$(GIT)build.Build=$(BUILD)\
				  -X=$(GIT)build.User=$(AUTHOR)"


build:
	go build -v $(LDFLAGS)
	
doc:
	swag init -g app/routes.go

docdeps:
	swag init --parseInternal --parseDependency -g app/routes.go

migrate-status:
	sql-migrate status -config=db/dbconfig.yml

migrate-up:
	sql-migrate up -config=db/dbconfig.yml

migrate-down:
	sql-migrate down -config=db/dbconfig.yml

migrate-new:
	sql-migrate new -config=db/dbconfig.yml

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store

.PHONY: build doc createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock
