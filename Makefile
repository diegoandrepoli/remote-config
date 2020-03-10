## Application makefile

start:
	go get -u ./...
	go run server.go


build:
	go get -u ./...
	go build
