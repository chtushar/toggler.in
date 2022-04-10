.DEFAULT_GOAL := build

check_migrate:
	which migrate || go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

add_migration: check_migrate
	 migrate create -dir=internal/db/migrations/ -seq -ext sql $(name)

check_lint:
	which staticcheck || go install honnef.co/go/tools/cmd/staticcheck@latest

check_sqlc:
	which sqlc || go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

lint: check_lint check_sqlc
	go vet ./...
	staticcheck ./...
	cd internal/db && sqlc compile

build:
	@sh -c './scripts/build.sh'

run: build
	./bin/app

install:
	./output/blog install

reset:
	./output/blog reset

gen:
	go generate ./...
