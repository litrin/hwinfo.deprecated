all: test readme

format:
	gofmt -w=true common

test: format
	golint common
#	go vet common
	go build common/*.go

readme:
	godoc2md github.com/mickep76/hwinfo >README.md
