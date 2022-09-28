.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/cmd ./cmd/main.go

clean:
	rm -rf ./bin
	rm -rf cover.out
	rm -rf cover.html

scrub: clean
	sls remove
	rm -rf ./serverless

deploy: clean build
	sls deploy --verbose

test: build
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
