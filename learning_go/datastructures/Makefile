APP_NAME := dsa
SOURCE := cmd/main.go

.PHONY: build
build:
	@go build -o ${APP_NAME} ${SOURCE}

.PHONY: run
run:
	@go run ${SOURCE}

.PHONY: test
test:
	@go test -v ./...

.PHONY: coverage
coverage:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

.PHONY: benchmark
benchmark:
	@go test -v -bench=. -benchmem

.PHONY: clean
clean:
	@rm -rf ${APP_NAME}
	@rm -rf coverage.out

.PHONY: help
help:
	@echo "make build - build the application"
	@echo "make run - run the application"
	@echo "make test - run the tests"
	@echo "make coverage - run the tests with coverage"
	@echo "make benchmark - run the benchmarks"
	@echo "make clean - remove the binary file"
	@echo "make help - display this help message"
