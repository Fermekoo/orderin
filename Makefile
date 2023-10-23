BINARY=orderin-api


# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/bin:$(PATH)
# Default Shell
export SHELL  := bash
# Type of OS: Linux or Darwin.
export OSTYPE := $(shell uname -s)

# # --- Tooling & Variables ----------------------------------------------------------------
include ./misc/make/tools.Makefile

install-deps: air  ## Install Development Dependencies (localy).
deps: $(AIR) ## Checks for Global Development Dependencies.
deps:
	@echo "Required Tools Are Available"


createmigrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "root:root@tcp(127.0.0.1:3307)/gokapster?multiStatements=true" -verbose up

migratedown:
	migrate -path db/migrations -database "root:root@tcp(127.0.0.1:3307)/gokapster?multiStatements=true" -verbose down

run:
	go run cmd/main.go

dev: $(AIR) ## Starts AIR ( Continuous Development app).
	air

test :
	go test -v -cover ./...

.PHONY: createmigrate migrateup migratedown run