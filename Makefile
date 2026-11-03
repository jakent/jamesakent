.PHONY: run test vet

run:
	go run ./cmd/site

test:
	go test ./...

vet:
	go vet ./...

.PHONY: acceptance
# In-process site by default; BASE_URL=http://localhost:8000 make acceptance drives a running stack.
acceptance:
	go test ./acceptance -count=1

.PHONY: up down acceptance-stack
up:
	docker compose up -d --build

down:
	docker compose down

# Full-stack run: includes the @stack scenarios (/repl, www redirect).
acceptance-stack:
	BASE_URL=http://localhost:$${PUBLIC_PORT:-8000} go test ./acceptance -count=1
