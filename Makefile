.PHONY: build
build:
	@echo "not ready yet"

.PHONY: clean
clean:
	@rm -rf ./bin

.PHONY: test
test:
	go test -v ./...
