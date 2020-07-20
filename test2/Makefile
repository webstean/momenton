#--- Makefile ----

# This how we want to name the binary output
BINARY=test2

HASH=`git rev-parse HEAD`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s -X main.Lastcommitsha=${HASH}"
# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}
    
# Installs our project: copies binaries
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
