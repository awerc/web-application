run: build
	./main -config=config.xml -log=log.txt
build:
	go build main.go
testAll:
	go test ./...