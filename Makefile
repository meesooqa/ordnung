.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: vet
lint:
	go vet ./...

.PHONY: test_race
test_race:
	go test -race -timeout=60s -count 1 ./...

.PHONY: test
test:
	go clean -testcache
	go test ./...
