.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test_race
test_race:
	go test -race -timeout=60s -count 1 ./...

.PHONY: test
test:
	go clean -testcache
	go test ./...

.PHONY: build
build:
	go build -o ./build/ordnung ./cmd/app/main.go
	cp .env build/
	mkdir -p build/var
	cp -r var/cache build/var/
