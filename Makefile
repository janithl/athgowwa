# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=athgowwa
BINARY_FOLDER=bin

all: test build
build: 
	$(GOBUILD) -o $(BINARY_FOLDER)/$(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...