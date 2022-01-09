build:
	go build main.go

run: build
	./api

watch:
	reflex -s -r '\.go$' make run
