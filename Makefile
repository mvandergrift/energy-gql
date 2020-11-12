.PHONY: build clean deploy
GIT_COMMIT := $(shell git rev-list -1 HEAD)

build:
	go build -ldflags="-s -w -X main.gitCommit=$(GIT_COMMIT)"

clean:
	rm -rf ./bin

test:
	go test ./...

generate:
	go generate ./...

run:
	go run server.go --sql

invoke:
	env GOOS=linux go build -o bin/energy-gql handlers/main.go
	sls invoke local -f query --path handlers/data.json

deploy: clean build
	env GOOS=linux go build -ldflags="-s -w -X main.gitCommit=$(GIT_COMMIT)" -o bin/energy-gql  handlers/main.go
	sls deploy
