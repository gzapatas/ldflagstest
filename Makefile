USER=$(shell id -u -n)
DATE=$(shell date)
VERSION=$(shell cat app.version)
LDFLAGS=-ldflags="-w -X 'main.Version=$(VERSION)' -X 'github.com/gzapatas/ldflagstest/build.Time=$(DATE)' -X 'github.com/gzapatas/ldflagstest/build.User=$(USER)'"

.PHONY : build run clean

build:
	go build $(LDFLAGS) -o app2

run:
	go run $(LDFLAGS) main.go

clean:
	go clean
	rm app