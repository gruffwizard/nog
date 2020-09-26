build:
        go env GOPATH
	go build -o target/nog main.go

run:
        go env GOPATH
	go run main.go

test:
        go env GOPATH
	go run main.go

lint:
        go env GOPATH
	go run main.go
