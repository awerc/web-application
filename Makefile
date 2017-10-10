run: build
	./main -config=config.xml -log=log.txt -xsd=xsd
build:
	go build main.go