.PHONY: build clean deploy commit

build: clean build/argh docs/index.xml

deploy: build commit

commit:
	git add -A
	git commit -m "rebuilding site `date`"
	git push origin master

docs/index.xml:
	cat feeds.txt | ./build/argh generate docs/index.xml

build/argh:
	go build -o build/argh .

clean:
	rm -f build/argh
	rm -f docs/index.xml
