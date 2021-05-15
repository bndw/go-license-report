.PHONY: build
build:
	go build -o ./bin/ .

.PHONY: test
test:
	go test -v ./...
