BINARY=gopt
.PRECIOUS: bin/${BINARY} dist/${BINARY}

all: bin/${BINARY}

bin/${BINARY}:
	go get
	go mod tidy
	mkdir -p bin
	go build -o $@

dist: dist/${BINARY}

dist/gopt: bin/${BINARY}
	mkdir -p dist
	cp -av $< $@
	strip $@
	upx -9 $@
	upx -t $@

install: dist/gopt
	test -e ~/common/bin && install -v $< ~/common/bin/${BINARY} || install -v $< ~/bin/${BINARY}

clean:
	rm -rf bin dist

