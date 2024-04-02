GO := go
CMD := cmd/

all: build

build: 
	cd $(CMD) && $(GO) build -o gd

.PHONY: all build