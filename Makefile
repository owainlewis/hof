all: codegen test

.PHONY: codegen
codegen:
	go run cmd/codegen/main.go

.PHONY: test
test:
	go test ./...

