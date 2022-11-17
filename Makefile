build:
	go build -o bin/testtools

build-docker:
	ko build ./internal/docker/ --local -t test

helm-dependency-build:
	 helm dependency build

deploy: helm-dependency-build
	helm upgrade --install testtools chart/testtools-container