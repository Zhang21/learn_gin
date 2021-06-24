.PHONY: build clean tool lint help

all: build

build:
	go build -v .

tool:
	go vet ./...; true
	gofmt .

lint:
	golint ./...

clean:
	rm -f learn_gin
	go clean -i .

help:
	@echo "make: compile package and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cache files and install archive"
