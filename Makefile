build:
	go build -o bin/kitly main.go

run:
	go run main.go

test:
	go test -v ./...
