.PHONY: build clean deploy commit docs

build: clean build/argh docs

commit:
	git add -A
	git commit -m "rebuilding site `date`"

docs:
	cat feeds.txt | go run main.go generate ./docs

build/argh:
	go build -o build/argh .

clean:
	rm -f build/argh
	rm -f docs/index.xml
