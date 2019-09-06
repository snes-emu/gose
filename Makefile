BINARY = gose
GOARCH = amd64

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${BRANCH}:${COMMIT}"
DEBUGLDFLAGS =-ldflags "-X main.VERSION=${BRANCH}:${COMMIT}:debug"

GOCMD = GO111MODULE=on go

# Build the project
all: build

.PHONY: build
build:
	${GOCMD} build ${LDFLAGS} -o ${BINARY} .

.PHONY: debug
debug:
	${GOCMD} build ${DEBUGLDFLAGS} -tags debug -o ${BINARY} .

.PHONY: linux
linux:
	GOOS=linux GOARCH=${GOARCH} ${GOCMD} build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH} .

.PHONY: macos
macos:
	GOOS=darwin GOARCH=${GOARCH} ${GOCMD} build ${LDFLAGS} -o ${BINARY}-macos-${GOARCH} .

.PHONY: windows
windows:
	GOOS=windows GOARCH=${GOARCH} ${GOCMD} build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe .

cross: linux macos windows

.PHONY: test
test:
	${GOCMD} get -t -v ./...; \
	${GOCMD} vet $$(${GOCMD} list ./... | grep -v /vendor/); \
	${GOCMD} test -v -race ./...; \

.PHONY: fmt
fmt:
	${GOCMD} fmt $$(${GOCMD} list ./... | grep -v /vendor/)
