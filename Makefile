GO ?= go
GOPATH := $(CURDIR)

game:
	@echo "----start build weiqi game server----"
	@pwd
	@rm -rf bin/server
	@export GOPATH=$(CURDIR)
	@go build src/main.go
	@mv main server
	@cp server ./bin/
	@rm -rf server
make_run:
	@make game
	@./bin/server