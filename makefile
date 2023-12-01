.PHONY: init
# init env
init:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
# lint code
lint:
	golangci-lint run -v

.PHONY: bench
# golang benchmark
benchmark:
	cd test && go test -bench=. -benchmem -benchtime=1s -run=none

.PHONY: build-cmd
# build pkg cmd
build-cmd:
	cd cmd && go build -o ../bin/mongo-cli
