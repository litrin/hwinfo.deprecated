all: test readme

format:
	gofmt -w=true common

test: format
	golint common
	golint cpu
#	go vet common
#	go build common/*.go
#	go build cpu/*.go

readme:
	godoc2md github.com/mickep76/hwinfo/common | grep -v Generated >README.md
	godoc2md github.com/mickep76/hwinfo/cpu | grep -v Generated >>README.md
