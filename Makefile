# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOGENERATE=$(GOCMD) generate
GOLIST=$(GOCMD) list

# Sources parameters
SOURCE_ENTRYPOINT=./benchdiff.go 

# Binary parameters
BINARY_NAME=benchdiff
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

# Tagets
build:
		$(GOBUILD) -o $(BINARY_PATH)
clean: 
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
test:
		$(GOTEST) -v -race -cover ./...