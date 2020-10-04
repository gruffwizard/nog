build:
	echo ${PWD}
	go build -o target/nog main.go
	target/nog

run:
	echo ${PWD}
	go env GOPATH
	go run main.go

test:
	echo ${PWD}
	go env GOPATH
	go run main.go

lint:
	echo ${PWD}
	go env GOPATH
	go run main.go

dummy-release:
	goreleaser --snapshot --skip-publish --rm-dist
