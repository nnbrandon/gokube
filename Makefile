PROJECT?=github.com/mrnguuyen/go_kube
APP?=advent
PORT?=8000
CONTAINER_IMAGE?=docker.io/herosekai/${APP}

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')


GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

test:
	go test -v -race ./...
