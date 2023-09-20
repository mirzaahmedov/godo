.DEFAULT_GOAL:=build

.PHONY=build run

build:
	@go build -o ./bin/httpserver ./cmd/httpserver

run: build
	@exec ./bin/httpserver
