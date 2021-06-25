.PHONY: build  clean help lint test tool

all: build

build:
	go build -v .

test:
	go test $$(go list ./... | grep -v 'vendor')

tool:
	go vet $$(go list ./... | grep -v 'vendor'); true
	gofmt .

lint:
	golint $$(go list ./... | grep -v 'vendor')

clean:
	rm -f learn_gin
	go clean -i .

help:
	@echo "make: compile package and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cache files and install archive"
