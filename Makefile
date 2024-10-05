.PHONY: build, run, clean
.DEFAULT_GOAL := all


build:
	go build

all: build
	@./go-pokedex-cli

clean:
	rm -rf ./go-pokedex-cli

