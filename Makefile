.PHONY: build clean

build: clean build/argh docs/index.xml

deploy: build
	git add -A
	git commit -m "rebuilding site `date`"

docs/index.xml:
	./build/argh generate docs/index.xml

build/argh:
	go build -o build/argh .

clean:
	rm -f build/argh

