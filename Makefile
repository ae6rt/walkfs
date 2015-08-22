
NAME := walkfs
ARCH := amd64

all: clean binaries 

binaries: 
	go build -o $(NAME)-darwin-$(ARCH)
	GOOS=linux GOARCH=$(ARCH) go build -o $(NAME)-linux-$(ARCH)

clean: 
	go clean
