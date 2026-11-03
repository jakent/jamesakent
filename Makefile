.PHONY: run test vet

run:
	go run ./cmd/site

test:
	go test ./...

vet:
	go vet ./...
