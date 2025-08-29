all: build

linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64  go build -o bin/linux/rettui ./rettui
	@echo "Done! Binary now is in bin/linux/rettui"

windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64  go build -o bin/windows/rettui.exe ./rettui
	@echo "Done! Binary now is in bin/windows/rettui.exe"

macos:
	@echo "Building for MacOS..."
	@GOOS=darwin GOARCH=amd64  go build -o bin/macos/rettui ./rettui
	@echo "Done! Binary now is in bin/macos/rettui"

build: linux windows macos

run:
	@go run ./rettui

clean:
	@echo "Cleaning bin/..."
	@rm -vf bin/*/*

.PHONY: all linux windows macos build run clean
