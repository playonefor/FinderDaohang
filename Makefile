GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
BINARY_NAME = FinderDaohang
CLI = adm
UNAME_S := $(shell uname -s)

all: serve

init:
	$(GOMOD) init $(module)

install:
	$(GOMOD) tidy

serve:
	$(GOCMD) run .

build:
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_NAME).exe -v ./
endif
ifeq ($(UNAME_S),Linux)
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -o  ./$(BINARY_NAME) -v ./
endif
ifeq ($(UNAME_S),Darwin)
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./$(BINARY_NAME) -v ./
endif

test: black-box-test user-acceptance-test

black-box-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainBlackBox
	make clean

user-acceptance-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainUserAcceptance
	make clean

ready-for-data:
	cp admin.db admin_test.db

clean:
	rm admin_test.db

.PHONY: all serve build test black-box-test user-acceptance-test ready-for-data clean
