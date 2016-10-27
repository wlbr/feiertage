
all: clean test build buildcmd
	
.PHONY: clean
clean: 
	rm -f feiertage 
	rm -rf bin/
	
run: 
	go run cmd/feiertage/command.go -region baden-württemberg 2016

test: 
	go test

build: feiertage.go region.go
	go build ./...
	
buildcmd: feiertage.go region.go cmd/feiertage/feiertage.go
	mkdir -p bin/linux_x86_64 bin/osx_x86_64 bin/win32 bin/win64
	GOOS=linux GOARCH=amd64   go build -o bin/linux_x86_64/feiertage cmd/feiertage/feiertage.go
	GOOS=darwin GOARCH=amd64  go build -o bin/osx_x86_64/feiertage   cmd/feiertage/feiertage.go
	GOOS=windows GOARCH=386   go build -o bin/win32/feiertage cmd/feiertage/feiertage.go
	GOOS=windows GOARCH=amd64 go build -o bin/win64/feiertage cmd/feiertage/feiertage.go

#install-go-deps:
	#go get -u ...
	
