
all: clean test build buildcmd
	
.PHONY: clean
clean: 
	rm -f feiertage

run: 
	go run cmd/feiertage/command.go -region baden-württemberg 2016

test: 
	go test

build: feiertage.go region.go
	go build ./...
	
buildcmd: feiertage.go region.go cmd/feiertage/feiertage.go
	go build cmd/feiertage/feiertage.go

#install-go-deps:
	#go get -u ...
	
