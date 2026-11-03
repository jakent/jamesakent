.PHONY: run test vet

run:
	go run ./cmd/site

test:
	go test ./...

vet:
	go vet ./...

.PHONY: acceptance
# In-process site by default; BASE_URL=http://localhost make acceptance drives a running stack.
acceptance:
	go test ./acceptance -count=1
