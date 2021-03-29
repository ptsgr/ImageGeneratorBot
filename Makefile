.PHONY: build
build:
	rm -rf build && mkdir build && go build -o build/imageGenerator -v ./cmd/imageGenerator

.PHONY: run
run:
	go run cmd/imageGenerator/main.go

.DEFAUL_GOAL := build