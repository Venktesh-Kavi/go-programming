APP_NAME := hello
SOURCE := helloworld/hello.go

all: build

.PHONY: build
build:
	@echo "building $(APP_NAME)"
	@go build -o $(APP_NAME) $(SOURCE)

.PHONY: run
run:
	@echo "running $(APP_NAME)"
	@go run $(SOURCE)

.PHONY: clean
clean:
	@rm -rf hello