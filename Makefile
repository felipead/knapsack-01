.PHONY: build
build:
	@go build -o bin/knapsack-01 cmd/main.go

.PHONY: clean
clean:
	@rm -rf ./bin

.PHONY: test
test:
	go test -v ./...
