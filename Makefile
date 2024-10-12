ready: vet lint format test

cover:
	go test -cover ./...

vet:
	go vet $(go list ./... | grep -v /vendor/)

tidy:
	go mod tidy

test:
	go test ./...

lint:
	golangci-lint run
	golint ./...

format:
	gofmt -d -l -e .

build:
	go build -o appservice ./cmd
