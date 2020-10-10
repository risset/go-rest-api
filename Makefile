GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN=go-rest-api

all: test build

build: 
	$(GOBUILD) -o $(BIN) -v

test: 

clean: 
	$(GOCLEAN)
	rm -f $(BIN)

run:
	$(GOBUILD) -o $(BIN) -v ./...
	./$(BIN)

deps:
	$(GOGET) github.com/lib/pq
	$(GOGET) github.com/go-chi/chi
	$(GOGET) github.com/go-chi/render
