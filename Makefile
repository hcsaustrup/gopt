GOOS ?= linux
GOARCH ?= amd64
BINARY = gopt
FULL_BINARY = ${BINARY}-${GOOS}-${GOARCH}

.PRECIOUS: bin/${FULL_BINARY} dist/${FULL_BINARY}

all: bin/${FULL_BINARY}

bin/${FULL_BINARY}:
	go get
	go mod tidy
	mkdir -p bin
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o $@

dist/${FULL_BINARY}:
	go get
	go mod tidy
	mkdir -p bin
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o $@ -ldflags '-s'
	upx -9 $@
	upx -t $@

dist: dist/${FULL_BINARY}

install: dist/${FULL_BINARY}
	test -e ~/common/bin && install -v $< ~/common/bin/${BINARY} || install -v $< ~/bin/${BINARY}

clean:
	rm -rf bin dist

