.PHONY: build
build:
	rm -rf build && mkdir build && CGO_ENABLED=0 go build -o build/imageGenerator -v ./cmd/imageGenerator
	docker build -t image_generator:latest .
	
.PHONY: run
run:
	go run cmd/imageGenerator/imageGenerator.go

.DEFAUL_GOAL := build