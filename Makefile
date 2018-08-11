BINARY = gose
GOARCH = amd64

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${BRANCH}:${COMMIT}"
DEBUGLDFLAGS =-ldflags "-X main.VERSION=${BRANCH}:${COMMIT}:debug"

# Build the project
all: build

.PHONY: build
build:
	go build ${LDFLAGS} -o ${BINARY} .

.PHONY: debug
debug:
	go build ${DEBUGLDFLAGS} -tags debug -o ${BINARY} .

.PHONY: linux
linux:
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH} .

.PHONY: macos
macos:
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-macos-${GOARCH} .

.PHONY: windows
windows:
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe .

cross: linux macos windows

.PHONY: test
test:
	go get -t -v ./...; \
    go vet $$(go list ./... | grep -v /vendor/); \
	go test -v -race ./...; \

.PHONY: fmt
fmt:
	go fmt $$(go list ./... | grep -v /vendor/)
