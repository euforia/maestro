
.test:
	go test -v -cover ./...

build:
	go get -d -v ./...
	go build ./...

install:
	go install ./...