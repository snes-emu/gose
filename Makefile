BINARY = gose
GOARCH = amd64

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${BRANCH}:${COMMIT}"

GOCMD = GO111MODULE=on go

# Build the project
all: build

.PHONY: build
build: deps
	${GOCMD} build ${LDFLAGS} -o ${BINARY} .

.PHONY: linux
linux: deps
	GOOS=linux GOARCH=${GOARCH} ${GOCMD} build ${LDFLAGS} -o ${BINARY}-linux-${GOARCH} .

.PHONY: macos
macos: deps
	GOOS=darwin GOARCH=${GOARCH} ${GOCMD} build ${LDFLAGS} -o ${BINARY}-macos-${GOARCH} .

.PHONY: windows
windows: deps
	GOOS=windows GOARCH=${GOARCH} ${GOCMD} build ${LDFLAGS} -o ${BINARY}-windows-${GOARCH}.exe .

.PHONY: cross
cross: linux macos windows

.PHONY: deps
deps:
	${GOCMD} get -v ./...

.PHONY: test
test: deps
	${GOCMD} vet $$(go list ./... | grep -v /vendor/); \
	${GOCMD} test -v -race ./...

.PHONY: fmt
fmt:
	${GOCMD} fmt $$(go list ./... | grep -v /vendor/)
