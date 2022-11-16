build:
	go build -o bin/testtools

build-docker:
	ko build ./internal/docker/ --local -t test