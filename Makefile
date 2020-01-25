# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOLIST=$(GOCMD) list
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt
BINARY_NAME=svd-statuspage
BINARY_NAME_LINUX64=$(BINARY_NAME)


all: test build
build:
				$(GOBUILD) -race -o $(BINARY_NAME) .
test:
				$(GOFMT) $(shell $(GOLIST) ./... | grep -v /vendor/)
				$(GOVET) $(shell $(GOLIST) ./... | grep -v /vendor/)
				$(GOTEST) -cover -race $(shell $(GOLIST) ./... | grep -v /vendor/)
clean:
				$(GOCLEAN)
				rm -f $(BINARY_NAME)
run: build
				./$(BINARY_NAME)

# Cross compilation
linux64:
				CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_LINUX64) -v .
