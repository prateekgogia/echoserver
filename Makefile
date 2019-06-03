IMAGE_NAME=echoserver

all: build test

.PHONY: build
build:
	docker build -t ${IMAGE_NAME} .

.PHONY: go-test
go-test:
	go test -v ./...

.PHONY: integration-test
integration-test: build
	./scripts/integrations_test.sh
	([ $$? -eq 0 ] && echo "Integrations Test PASSED!") || echo "FAILURE!"

.PHONY: test
test: go-test integration-test

.PHONY: run-server
run-server: build
	docker run -d --rm --name=server echoserver /bin/server
