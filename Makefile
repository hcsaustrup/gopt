
all: bin/gopt

bin/gopt:
	go get
	go mod tidy
	mkdir -p bin
	go build -o $@

bin/gopt-compressed: bin/gopt
	cp -av $< $@
	strip $@
	upx -9 $@
	upx -t $@

install: bin/gopt-compressed
	test -e ~/common/bin && install -v $< ~/common/bin/ || install -v $< ~/bin

clean:
	rm -rf bin

