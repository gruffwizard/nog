build:
	echo ${PWD}
	go env GOPATH
	go get -u ./...
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
