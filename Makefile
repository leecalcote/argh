.PHONY: build clean deploy commit

build: clean build/argh docs/index.xml

commit:
	git add -A
	git commit -m "rebuilding site `date`"

docs/index.xml:
	cat feeds.txt | ./build/argh generate docs/index.xml

build/argh:
	go build -o build/argh .

clean:
	rm -f build/argh
	rm -f docs/index.xml
