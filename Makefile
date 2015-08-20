PACKAGE:=$(shell git config --get remote.origin.url | sed -e 's!^git@!!' -e 's!\.git$!!' -e 's!:!/!')

all: test readme

format:
	gofmt -w=true .

test: format
	golint .
	go vet .
	go build

readme:
	godoc2md ${PACKAGE} >README.md
