.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run ./...
.PHONY:lint

vet: lint
	go vet ./...
	shadow ./...
.PHONY:vet

build: vet
	go build quiz.go
.PHONY:build

run: build
	./quiz
.PHONY:run

clean:
	rm quiz
.PHONY:clean