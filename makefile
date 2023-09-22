
.PHONY: install
install:
	go mod tidy

.PHONY: test
test: install
	go test ./... -race

.PHONY: run
run:
	go run ./...