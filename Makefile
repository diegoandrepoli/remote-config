## Application makefile

start:
	go get -u ./...
	go run server.go


build:
	go build
