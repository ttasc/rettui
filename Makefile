all: build

build:
	@echo "Building..."
	@go build -o bin/rettui ./src
	@echo "Done! Binary now is in bin/rettui"

run:
	@go run ./src

clean:
	@echo "Cleaning bin/..."
	@rm -vf bin/*

.PHONY: all build run clean
