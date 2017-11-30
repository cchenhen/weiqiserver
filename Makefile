GO ?= go
GOPATH := $(CURDIR)

game:
	@echo "----start build weiqi game server----"
	@pwd
	@rm -rf server 
	@rm -rf main
	@export GOPATH=$(CURDIR)
	@go build src/main.go
	@mv main server