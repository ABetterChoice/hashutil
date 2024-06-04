all: .format

.PHONY: .format

.format:
	go mod tidy
	gofmt -w .
	goimports -w .
