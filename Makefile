build:
	go build -o app main.go

start:
	go run main.go

tests:
	go test -v ./test/...
