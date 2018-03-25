SOURCES = $(shell find $(SOURCEDIR) -name '*.go')
BINARY=event

.PHONY: build clean

build:
	go build -o ${BINARY} $(SOURCES)

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi