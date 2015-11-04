
all: clean test build 
	
.PHONY: clean
clean: 
	rm -f feiertage

run: 
	go run feiertage.go

test: 
	go test

build: feiertage.go
	go build feiertage.go

#install-go-deps:
	#go get -u ...
	
