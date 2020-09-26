build:
	echo ${PWD}
	go env GOPATH
	go build -o target/nog main.go

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
