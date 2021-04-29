UNFORMATTED := $(shell gofmt -l . )
BUILD_SERVER := $(shell cd server/ && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/app -v ./...)
ci:
	[ -z ${UNFORMATTED} ] && exit 0
	find . -name go.mod -execdir go test ./... \;

build-server:
	${BUILD_SERVER}

deploy:
	exec ${BUILD_SERVER} && (cd server; bin/app)
