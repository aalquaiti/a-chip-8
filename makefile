.PHONY: help create activate install
.DEFAULT: help

help:
	@echo "    build      Compiles app"
	@echo "    app    Run app"

build:
	@go build -o bin/chip8 ./cmd/chip8/main.go

run:
	@go run ./cmd/chip8/main.go
