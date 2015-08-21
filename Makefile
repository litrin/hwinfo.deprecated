all: test readme

test:
	cd common ;\
	make
	cd cpu ;\
	make
	cd mem ;\
	make

readme:
	godoc2md github.com/mickep76/hwinfo/common | grep -v Generated >README.md
	godoc2md github.com/mickep76/hwinfo/cpu | grep -v Generated >>README.md
