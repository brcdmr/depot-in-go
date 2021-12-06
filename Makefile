SOURCES := $(wildcard *go cmd/*/*.go pkg/*/*.go)

VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

## we must have tagged the repo at least once for VERSION to work
ifeq ($(VERSION),)
	VERSION = 0.0.1
endif

depot: $(SOURCES)
	go build -ldflags "-X main.version=${VERSION}" -o $@ ./cmd/depot

.PHONY: lint
lint:
	golangci-lint run

.PHONY: committed
committed:
	@git diff --exit-code > /dev/null || (echo "** COMMIT YOUR CHANGES FIRST **"; exit 1)

docker: $(SOURCES) build/Dockerfile
	docker build -t depot-in-go:latest . -f build/Dockerfile --build-arg VERSION=$(VERSION)

.PHONY: publish
publish: committed lint
	make docker
	docker tag  depot-in-go:latest brcdmr/depot-in-go:$(VERSION)
	docker push brcdmr/depot-in-go:$(VERSION)
