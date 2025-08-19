all: build

build:
	@echo "Building..."
	@go build -o bin/restt ./src
	@echo "Done! Binary now is in bin/restt"

run:
	@go run ./src

clean:
	@echo "Cleaning..."
	@rm -vf bin/restt

.PHONY: all build run clean
