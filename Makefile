GO_CMD = go
all: build test
test:
	$(GO_CMD) test -v ./...
build:
	$(GO_CMD) build -o main ./main.go
clean:
	rm -f main