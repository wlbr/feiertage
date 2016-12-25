
all: clean test build buildcmd

.PHONY: clean
clean:
	rm -f feiertage
	rm -rf bin/

run:
	go run cmd/feiertage/feiertage.go -region baden-württemberg 2016

debug:
	dlv debug cmd/feiertage/feiertage.go -- -region baden-württemberg 2016


test:
	go test

build: feiertage.go region.go
	go build ./...

buildcmd: feiertage.go region.go cmd/feiertage/feiertage.go
	mkdir -p bin
	GOOS=linux GOARCH=amd64   go build -o bin/feiertage_linux_x86_64 cmd/feiertage/feiertage.go
	GOOS=darwin GOARCH=amd64  go build -o bin/feiertage_osx_x86_64   cmd/feiertage/feiertage.go
	GOOS=windows GOARCH=386   go build -o bin/feiertage_win32        cmd/feiertage/feiertage.go
	GOOS=windows GOARCH=amd64 go build -o bin/feiertage_win64        cmd/feiertage/feiertage.go


