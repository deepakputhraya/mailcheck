UNFORMATTED := $(shell gofmt -l . )

ci:
	[ -z ${UNFORMATTED} ] && exit 0
	find . -name go.mod -execdir go test ./... \;

deploy:
	cd server/ && go build -o bin/app -v ./... && bin/app

