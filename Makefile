.PHONY: build clean

build: clean build/argh

docs/index.xml:
	./build/argh generate

build/argh:
	go build -o build/argh .

clean:
	rm -f build/argh
