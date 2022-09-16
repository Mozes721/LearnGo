.DEFAULT_GOAL := main.out
fmt:
	go fmt .
.PHONY:fmt
lint:fmt	
	golint .

vet: fmt
	go vet .
.PHONY:vet
build: vet
	go build main.go
.PHONY:build
